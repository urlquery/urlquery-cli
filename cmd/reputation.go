package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/urlquery/urlquery-cli/internal/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reputationCmd = &cobra.Command{
	Use:   "reputation <url>",
	Short: "Check the reputation of a URL.",
	Long: `Check the reputation of a given URL using urlquery.net

This command queries the urlquery reputation API.

Requires a valid API key (set via 'config set apikey <value>' or the --apikey flag).

Example:
	urlquery-cli reputation http://example.com
	urlquery-cli reputation www.youtube.com/watch?v=dQw4w9WgXcQ
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reputation_url := args[0]

		// API key
		apikey := viper.GetString("apikey")

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

		summary := viper.GetBool("summary")
		if summary {

			fmt.Println("🔎 Reputation Summary")
			fmt.Println("────────────────────────────────────────────────────────────")
			fmt.Printf("🔗 URL:     %s\n", response.Url)

			// Optionally add a verdict icon
			var verdictIcon string
			switch response.Verdict {
			case "malicious":
				verdictIcon = "🚫"
			case "suspicious":
				verdictIcon = "⚠️"
			case "benign":
				verdictIcon = "✅"
			default:
				verdictIcon = "❔"
				verdictIcon = ""
			}
			fmt.Printf("🛡️  Verdict: %s %s\n", verdictIcon, strings.Title(response.Verdict))
			return
		}

		// Default JSON output
		out, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error formatting response:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))

	},
}
