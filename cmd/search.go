package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urlquery/urlquery-cli/internal/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var limitSearch int
var offsetSearch int

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for reports in urlquery.net.",
	Long: `Searches urlquery.net for reports related to a domain, IP, or keyword.
For more details check out: https://urlquery.net/help/search`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search_query := args[0]

		// API key
		apikey := viper.GetString("apikey")

		// Initialize API client
		client, err := api.NewClient(api.ApiKey(apikey))
		if err != nil {
			fmt.Println("Failed to create API client:", err)
			os.Exit(1)
		}

		// Perform search request
		results, err := client.Search(search_query, limitSearch, offsetSearch)
		if err != nil {
			fmt.Printf("Error searching reports: %v\n", err)
			os.Exit(1)
		}

		summary := viper.GetBool("summary")
		if summary {

			fmt.Printf("🔍 Search Query: %s\n", results.Query)
			fmt.Printf("Hits:    %d\n", results.TotalHits)
			fmt.Printf("Limit:   %d\n", results.Limit)
			fmt.Printf("Offset:  %d\n\n", results.Offset)

			for _, v := range results.Reports {
				url := v.Url.Addr
				if len(url) > 76 {
					url = url[:70] + " (...)"
				}

				fmt.Println("\n--------------------------------------------------------------------------------")
				// fmt.Printf("📝 Report ID:  %s\n", v.ID)
				color.New(color.Bold).Printf("📝 Report ID:  %s\n", v.ID)
				fmt.Printf("🔗 URL:        %s\n", url)
				fmt.Printf("🚨 Detections: %d\n", v.Stats.AlertCount.Urlquery)
				fmt.Printf("🏷️  Tags:       %s\n", strings.Join(v.Tags, " "))
			}
			fmt.Println("")
			return
		}

		// Default full JSON output
		output, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error formatting response:", err)
			os.Exit(1)
		}
		fmt.Println(string(output))

	},
}
