package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Use title case. Any combination of lower/uppercase will map to this value
// when decoding into a struct
type Joke struct {
	ID        string `json:"id"`
	Punchline string `json:"joke"`
	Status    int    `json:"status"`
}

// dadjokeCmd represents the dadjoke command
var dadjokeCmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "Prints a random dadjoke",
	Long:  `Prints a random dadjoke from powered by the https://icanhazdadjoke.com/api`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return randomDadJoke("https://icanhazdadjoke.com", os.Stdout)
	},
}

func randomDadJoke(baseURL string, out io.Writer) error {
	request, err := http.NewRequest(
		http.MethodGet,
		baseURL,
		nil)
	if err != nil {
		// %w wraps the error, so it can be later unwrapped with errors.Unwrap
		return fmt.Errorf("Failed to make a dadjoke request object: %w", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "🔥cli https://github.com/gwwar/firecli")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Failed to make request to dadjoke api: %w", err)
	}
	joke := &Joke{}
	err = json.NewDecoder(response.Body).Decode(joke)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json dadjoke: %w", err)
	}
	if joke.Punchline == "" {
		return fmt.Errorf("Got a joke response with %v", joke)
	}
	_, err = fmt.Fprintln(out, joke.Punchline)
	return err
}

func init() {
	rootCmd.AddCommand(dadjokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dadjokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dadjokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
