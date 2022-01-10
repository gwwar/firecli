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
		width := 80
		writer := newWriter(width)
		hideCat, e := cmd.Flags().GetBool("textonly")
		if e != nil {
			return e
		}
		showCat := !hideCat
		if isInputFromPipe() {
			return catsay(os.Stdin, writer, showCat)
		}
		if len(args) > 0 {
			data := []byte(args[0])
			return catsay(bytes.NewReader(data), writer, showCat)
		}
		message, e := cmd.Flags().GetString("message")
		if e != nil {
			return e
		}
		data := []byte(message)
		return catsay(bytes.NewReader(data), writer, showCat)
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

func writeWord(w *wordWrappedWriter, word string) error {
	var err error
	var totalWritten int

	if w.lineWidth == w.columnsLeft {
		totalWritten, err = fmt.Print(w.dleft, " ")
		w.columnsLeft -= totalWritten
	}

	columns := utf8.RuneCountInString(word) + 1

	if w.columnsLeft > columns {
		totalWritten, err = fmt.Print(word, " ")
		w.columnsLeft -= totalWritten
	} else if columns > w.lineWidth {
		//word is larger than linewidth, split it
		runes := []rune(word)
		first := string(runes[0 : w.columnsLeft-1])
		rest := string(runes[w.columnsLeft-1 : columns])
		_, err = fmt.Println(first + "-")
		w.columnsLeft = w.lineWidth
		if err != nil {
			return err
		}
		return writeWord(w, rest)
	} else {
		//add a newline
		if w.columnsLeft > 2 {
			_, err = fmt.Println(strings.Repeat(" ", w.columnsLeft-2), w.dright)
			if err != nil {
				return err
			}
		} else if w.columnsLeft == 2 {
			_, err = fmt.Println("", w.dright)
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Println(w.dright)
			if err != nil {
				return err
			}
		}
		w.columnsLeft = w.lineWidth
		return writeWord(w, word)
	}
	return err
}

func catsay(r io.Reader, w *wordWrappedWriter, showCat bool) error {
	//draw top
	fmt.Println(strings.Repeat(w.dtop, w.lineWidth))
	// message
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			e := writeWord(w, word)
			if e != nil {
				return e
			}
		}
	}
	if w.columnsLeft != w.lineWidth {
		if w.columnsLeft > 2 {
			_, e := fmt.Println(strings.Repeat(" ", w.columnsLeft-2), w.dright)
			if e != nil {
				return e
			}
		} else if w.columnsLeft == 2 {
			_, e := fmt.Println("", w.dright)
			if e != nil {
				return e
			}
		} else {
			_, e := fmt.Println(w.dright)
			if e != nil {
				return e
			}
		}
	}
	//draw bottom
	fmt.Println(strings.Repeat(w.dbottom, w.lineWidth))
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
