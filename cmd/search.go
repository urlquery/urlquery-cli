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

var limitSearch int
var offsetSearch int

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for reports in urlquery.net.",
	Long:  "Searches urlquery.net for reports related to a domain, IP, or keyword.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search_query := args[0]

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

		// Perform search request
		results, err := client.Search(search_query, limitSearch, offsetSearch)
		if err != nil {
			fmt.Printf("Error searching reports: %v\n", err)
			os.Exit(1)
		}

		format := "json"
		switch format {
		case "summary":
			fmt.Println("Search Query: ", results.Query)

			for _, v := range results.Reports {
				fmt.Println("")
				fmt.Printf("Report ID: %s\n", v.ID)
				fmt.Printf("URL:  %s\n", v.Url.Addr)
				fmt.Printf("Tags: %s\n", strings.Join(v.Tags, ","))
				fmt.Printf("Detections: %d\n", v.Stats.AlertCount.Urlquery)
			}

		case "json":
			fallthrough
		default:
			output, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Println("Error formatting response:", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}

	},
}
