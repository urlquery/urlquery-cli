package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/urlquery/urlquery-api-go"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var submitCmd = &cobra.Command{
	Use:   "submit <url>",
	Short: "Submit a URL for sandbox analysis and threat detection.",
	Long: `Submit a URL to urlquery.net for sandbox analysis and threat detection.

You can customize the submission with:
  - access: 	public, restricted, or private
  - useragent: 	override the default browser user-agent
  - tags: 		comma-separated values to label the submission

Requires an API key (set via 'config set apikey <value>' or --apikey).

Example:
  urlquery-cli submit https://example.com
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		submit_url := args[0]

		apikey := viper.GetString("apikey")
		if apikey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
			os.Exit(1)
		}

		job := urlquery.SubmitJob{
			Url: submit_url,
		}

		job.UserAgent = viper.GetString("useragent")
		if job.UserAgent == "" {
			job.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:138.0) Gecko/20100101 Firefox/138.0"
		}

		// Validate tags
		tags := viper.GetString("tags")
		if tags != "" {
			tmpTags := strings.Split(tags, ",")
			var validTags []string
			for _, tag := range tmpTags {
				if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(strings.Trim(tag, " ")) {
					fmt.Printf("Removed invalid tag: %s (tags must be alphanumeric or underscore)\n", tag)
					continue
				}
				validTags = append(validTags, strings.Trim(tag, " "))
			}
			job.Tags = validTags
		}

		// Validate access value
		access := viper.GetString("access")
		validAccess := map[string]bool{
			"public":     true,
			"restricted": true,
			"private":    true,
		}
		if !validAccess[access] {
			access = "public"
		}
		job.Access = access

		// Submit URL
		client := urlquery.NewClient(apikey)
		client.Submit(job)

		response, err := client.Submit(job)
		if err != nil {
			fmt.Printf("Error querying URL: %v\n", err)
			return
		}

		summary := viper.GetBool("summary")
		if summary {

			bold := color.New(color.Bold).SprintFunc()
			fmt.Println("Submitted URL:")
			fmt.Printf("🔗 URL:      %s\n", bold(response.Url.Addr))
			fmt.Printf("🆔 Queue ID: %s\n", response.QueueID)
			fmt.Printf("📊 Status:   %s\n", response.Status)
			fmt.Println("")
			fmt.Printf("https://urlquery.net/queue/%s\n", response.QueueID)
			return
		}

		// Default JSON output
		output, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error formatting response:", err)
			os.Exit(1)
		}
		fmt.Println(string(output))

	},
}

// Sub command of submit which returns the current status of a submission
var submitStatusCmd = &cobra.Command{
	Use:   "status <queue_id>",
	Short: "View the current processing status of a submitted URL.",
	Long: `Check the analysis status of a previously submitted URL using its queue ID.

The queue ID is returned when a URL is submitted. Use this command to poll its current state.

Possible statuses:
  - queued     : Waiting in line for analysis
  - processing : Sandbox is actively analyzing the URL
  - analyzing  : Data is currently being analysed and finalised
  - done       : Analysis is complete and results are available

Requires an API key (set via 'config set apikey <value>' or --apikey).

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

		client := urlquery.NewClient(apikey)

		response, err := client.GetQueueStatus(queue_id)
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
