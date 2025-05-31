package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urlquery/urlquery-cli/internal/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var submitCmd = &cobra.Command{
	Use:   "submit <url>",
	Short: "Submit a URL for analysis.",
	Long: `Submits a URL to urlquery.net for analysis and threat detection.

The URL will be analyzed in a sandbox environment. The submission can be made with custom access visibility
(public, restricted, private) and optional user-agent settings.

Requires an API key (set via 'config set apikey <value>' or --apikey).

Example:
	urlquery-cli submit http://example.com
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		submit_url := args[0]

		apikey := viper.GetString("apikey")
		if apikey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
			os.Exit(1)
		}

		var job api.SubmitJob
		job.Url = submit_url

		job.UserAgent = viper.GetString("useragent")
		if job.UserAgent == "" {
			job.UserAgent = "default"
		}

		job.Access = viper.GetString("access")
		if job.Access == "" {
			job.Access = "public"
		}

		client, err := api.NewClient(api.ApiKey(apikey))
		if err != nil {
			fmt.Println("Failed to create API client:", err)
			os.Exit(1)
		}

		response, err := client.Submit(job)
		if err != nil {
			fmt.Printf("Error querying URL: %v\n", err)
			return
		}

		format := "json"
		switch format {
		case "summary":
			fmt.Println("Submitted URL:")
			fmt.Printf("URL: %s\n", response.Url.Addr)
			fmt.Printf("Queue ID: %s\n", response.QueueID)
			fmt.Printf("Status: %s\n", response.Status)

		case "json":
			fallthrough
		default:
			output, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				fmt.Println("Error formatting response:", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

// Sub command of submit which returns the current status of a submission
var submitStatusCmd = &cobra.Command{
	Use:   "status <queue_id>",
	Short: "Check the status of a submitted URL.",
	Long: `Checks the processing status of a submitted URL using its queue ID.

The queue ID is returned after submission and can be used to poll for analysis status.

Example:
	urlquery-cli submit status 902d9135-12fe-4e75-95bb-a6d1e8c79ed1
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queue_id := args[0]

		apikey := viper.GetString("apikey")
		if apikey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
			os.Exit(1)
		}

		client, err := api.NewClient(api.ApiKey(apikey))
		if err != nil {
			fmt.Println("Failed to create API client:", err)
			os.Exit(1)
		}

		response, err := client.QueueStatus(queue_id)
		if err != nil {
			fmt.Printf("Error fetching status for Queue ID %s: %v\n", queue_id, err)
			os.Exit(1)
		}

		output, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error formatting response:", err)
			os.Exit(1)
		}

		fmt.Println(string(output))
	},
}
