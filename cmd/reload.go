package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// reloadCmd represents the reload command
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reloads the prometheus configuration",
	Long:  `If changes were made to the prometheus configuration, this triggers a reload without needing to restart the service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return reload("http://localhost:9000", os.Stdout)
	},
}

func reload(prometheusBaseURL string, out io.Writer) error {
	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/-/reload", prometheusBaseURL),
		nil)
	if err != nil {
		// %w wraps the error, so it can be later unwrapped with errors.Unwrap
		return fmt.Errorf("Failed to make a reload request object: %w", err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("Failed to make a reload request to prometheus: %w", err)
	}
	if response.StatusCode == 200 {
		fmt.Fprintln(out, "Successfully reloaded prometheus configs, visit", prometheusBaseURL)
	} else {
		fmt.Fprintln(out, "Failed to reload prometheus configs")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(reloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
