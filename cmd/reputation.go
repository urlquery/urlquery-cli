package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urlquery-cli/internal/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reputationCmd = &cobra.Command{
	Use:   "reputation <url>",
	Short: "Check the reputation of a URL.",
	Long: `Check the reputation of a given URL using urlquery.net.

This command queries the urlquery reputation API.

Requires a valid API key (set via 'config set apikey <value>' or the --apikey flag).

Example:
	urlquery-cli reputation http://example.com
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reputation_url := args[0]

		if args[0] == "help" {
			cmd.Help()
			return
		}

		// Ensure API key is available
		apikey := viper.GetString("apikey")
		if apikey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
			os.Exit(1)
		}

		// Initialize API client
		client, err := api.NewClient(api.ApiKey(apikey))
		if err != nil {
			fmt.Println("Failed to create API client:", err)
			os.Exit(1)
		}

		// Fetch reputation data
		response, err := client.CheckReputation(reputation_url)
		if err != nil {
			fmt.Printf("Error querying URL reputation: %v\n", err)
			os.Exit(1)
		}

		format := "json"
		switch format {
		case "summary":
			fmt.Println("Reputation Summary:")
			fmt.Printf("  URL:     %s\n", response.Url)
			fmt.Printf("  Verdict: %s\n", response.Verdict)
		case "json":
			fallthrough
		default:
			out, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				fmt.Println("Error formatting response:", err)
				os.Exit(1)
			}
			fmt.Println(string(out))
		}

	},
}
