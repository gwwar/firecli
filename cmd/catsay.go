package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

// catsayCmd represents the catsay command
var catsayCmd = &cobra.Command{
	Use:   "catsay",
	Short: "A speaking cat",
	Long: `Catsay generates an ASCII picture of a cat saying something provided by the user.

If run with no arguments, it prints a default message. This command also accepts piped input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		out := os.Stdout
		width := 80
		writer := newWriter(width)
		hideCat, e := cmd.Flags().GetBool("textonly")
		if e != nil {
			return e
		}
		showCat := !hideCat
		if isInputFromPipe() {
			return catsay(out, os.Stdin, writer, showCat)
		}
		if len(args) > 0 {
			data := []byte(args[0])
			return catsay(out, bytes.NewReader(data), writer, showCat)
		}
		message, e := cmd.Flags().GetString("message")
		if e != nil {
			return e
		}
		data := []byte(message)
		return catsay(out, bytes.NewReader(data), writer, showCat)
	},
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

type wordWrappedWriter struct {
	Writer      io.Writer
	lineWidth   int
	columnsLeft int
	dtop        string
	dleft       string
	dright      string
	dbottom     string
}

func newWriter(width int) *wordWrappedWriter {
	writer := &wordWrappedWriter{
		lineWidth:   width,
		columnsLeft: width,
		dtop:        "-",
		dleft:       ".",
		dright:      ".",
		dbottom:     "-",
	}
	return writer
}

func writeWord(out io.Writer, w *wordWrappedWriter, word string) error {
	var err error

	if w.lineWidth == w.columnsLeft {
		_, err = fmt.Fprint(out, w.dleft, " ")
		if err != nil {
			return err
		}
	}

	columns := utf8.RuneCountInString(word)

	if w.columnsLeft >= columns {
		_, err = fmt.Fprint(out, word)
		w.columnsLeft -= columns
	} else if columns > w.lineWidth {
		//word is larger than linewidth, split it
		runes := []rune(word)
		first := string(runes[0 : w.columnsLeft-1])
		rest := string(runes[w.columnsLeft-1 : columns])
		_, err = fmt.Fprintln(out, first+"-")
		w.columnsLeft = w.lineWidth
		if err != nil {
			return err
		}
		return writeWord(out, w, rest)
	} else {
		//add a newline
		e := lineReturn(out, w)
		if e != nil {
			return e
		}
		return writeWord(out, w, strings.TrimPrefix(word, " "))
	}
	return err
}

func lineReturn(out io.Writer, w *wordWrappedWriter) error {
	spaces := strings.Repeat(" ", w.columnsLeft)
	_, e := fmt.Fprintln(out, spaces, w.dright)
	if e != nil {
		return e
	}
	w.columnsLeft = w.lineWidth
	return nil
}

func catsay(out io.Writer, r io.Reader, w *wordWrappedWriter, showCat bool) error {
	//draw top
	fmt.Fprintln(out, strings.Repeat(w.dtop, w.lineWidth+4))
	// message
	scanner := bufio.NewScanner(bufio.NewReader(r))
	var spaceAndWord string
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			if w.columnsLeft == w.lineWidth {
				spaceAndWord = word
			} else {
				spaceAndWord = " " + word
			}
			e := writeWord(out, w, spaceAndWord)
			if e != nil {
				return e
			}
		}
	}
	e := lineReturn(out, w)
	if e != nil {
		return e
	}
	//draw bottom
	fmt.Println(strings.Repeat(w.dbottom, w.lineWidth+4))
	// ascii cat
	if showCat {
		dat, e := os.ReadFile("assets/cat.text")
		if e != nil {
			return fmt.Errorf("couldn't read catsay asset: %w", e)
		}
		_, e = fmt.Fprintln(os.Stdout, string(dat))
		return e
	}
	return nil
}

func init() {
	rootCmd.AddCommand(catsayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	catsayCmd.PersistentFlags().StringP("message", "m", "Hello world!", "Message to print")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	catsayCmd.Flags().BoolP("textonly", "t", false, "Prints text without the ascii cat")
}
