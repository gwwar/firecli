package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// catsayCmd represents the catsay command
var catsayCmd = &cobra.Command{
	Use:   "catsay",
	Short: "A speaking cat",
	Long: `Catsay generates an ASCII picture of a cat saying something provided by the user.

	If run with no arguments, it accepts standard input, word-wraps the message given, and prints the cat saying the
	given message on standard output`,
	RunE: func(cmd *cobra.Command, args []string) error {

		dat, err := os.ReadFile("assets/cat.text")
		if ( err != nil ) {
			return fmt.Errorf("couldn't read catsay asset: %w", err)
		}
		fmt.Println(string(dat))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(catsayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catsayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catsayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
