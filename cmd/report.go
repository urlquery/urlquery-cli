package cmd

import (
	"fmt"
	"os"

	"github.com/urlquery-cli/internal/api"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reportCmd = &cobra.Command{
	Use:   "report <report_id> <report|screenshot|domain_graph|resource> [hash]",
	Short: "Fetch report details or download artifacts",
	Long: `Retrieve data from a submitted URL scan by its Report ID.

You can download the full JSON report, screenshot, domain graph visualization, or a specific resource file (by its hash).
  - report        JSON report with scan metadata and results
  - screenshot    Screenshot of the loaded URL
  - domain_graph  Visual representation of domain relationships
  - resource      Specific resource from the scan (hash)

All downloaded files are saved in the output directory (default: current directory, or use --output).

Usage:
  urlquery-cli report <report_id> report
  urlquery-cli report <report_id> screenshot
  urlquery-cli report <report_id> domain_graph
  urlquery-cli report <report_id> resource <hash>

Examples:
  urlquery-cli report 0bd0eba2-82c7-4c39-bdec-dfa9d9fc582e report
  urlquery-cli report 0bd0eba2-82c7-4c39-bdec-dfa9d9fc582e screenshot
  urlquery-cli report 0bd0eba2-82c7-4c39-bdec-dfa9d9fc582e resource 4f9d4b...`,
	Args: cobra.MinimumNArgs(2), // At least 2 arguments required

	Run: func(cmd *cobra.Command, args []string) {

		report_id := args[0]
		action := args[1]

		// Validate Report UUID
		if _, err := uuid.Parse(report_id); err != nil {
			fmt.Printf("Error: '%s' is not a valid UUID.\n", report_id)
			os.Exit(1)
		}

		// Check API key
		apikey := viper.GetString("apikey")
		if apikey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
			os.Exit(1)
		}

		client, err := api.NewClient(api.ApiKey(apikey))
		if err != nil {
			fmt.Println("Failed", err)
			os.Exit(1)
		}

		// Handle report data
		validActions := map[string]bool{
			"report":       true,
			"screenshot":   true,
			"domain_graph": true,
			"resource":     true,
		}
		if validActions[action] {
			fmt.Printf("Fetching %s for Report: %s\n", action, report_id)

			// Fetch Report
			if action == "report" {
				report, err := client.GetReport(report_id)
				if err != nil {
					fmt.Println("Failed", err)
					os.Exit(1)
				}

				report_filename := fmt.Sprintf("report_%s.json", report_id)
				f, _ := os.OpenFile(report_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

				fmt.Println(len(report.Bytes()))

				f.Write(report.Bytes())
				f.Close()
			}

			// Fetch Domain Graph
			if action == "domain_graph" {
				domain_graph_filename := fmt.Sprintf("domain_graph_%s.gif", report_id)

				data_domain_graph, _ := client.GetDomainGraph(report_id)
				f, _ := os.OpenFile(domain_graph_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

				f.Write(data_domain_graph)
				f.Close()
			}

			// Fetch Screenshot
			if action == "screenshot" {
				screenshot_filename := fmt.Sprintf("screenshot_%s.png", report_id)

				data_screenshot, _ := client.GetScreenshot(report_id)
				f, _ := os.OpenFile(screenshot_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

				f.Write(data_screenshot)
				f.Close()
			}

			// Handle downloading a resource
			if action == "resource" {
				if len(args) < 3 {
					fmt.Println("Error: Missing resource hash.\nUsage: urlquery-cli report <report_id> resource <hash>")
					os.Exit(1)
				}
				hash := args[2]

				fmt.Printf("Downloading resource %s from Report ID: %s\n", hash, report_id)
				// Call API to download the resource
				return
			}

		}

	},
}
