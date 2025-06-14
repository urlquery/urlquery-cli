package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urlquery/urlquery-api-go"
	"github.com/urlquery/urlquery-cli/internal/api"

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

All downloaded files are saved in the output directory (default: current directory, or set via 'config set output <calue>' use --output).

Usage:
  urlquery-cli report <report_id> report
  urlquery-cli report <report_id> screenshot
  urlquery-cli report <report_id> domain_graph
  urlquery-cli report <report_id> resource <hash>

Examples:
  urlquery-cli report 82c4121d-d037-4d60-9f74-517bf00091ce report
  urlquery-cli report 82c4121d-d037-4d60-9f74-517bf00091ce screenshot
  urlquery-cli report 82c4121d-d037-4d60-9f74-517bf00091ce resource 4f9d4b...`,
	Args: cobra.MinimumNArgs(2), // At least 2 arguments required

	Run: func(cmd *cobra.Command, args []string) {

		report_id := args[0]
		action := args[1]

		output_directory := filepath.Clean(viper.GetString("output")) + string(os.PathSeparator)

		// Validate Report UUID
		if _, err := uuid.Parse(report_id); err != nil {
			fmt.Printf("Error: '%s' is not a valid UUID.\n", report_id)
			os.Exit(1)
		}

		apikey := viper.GetString("apikey")

		client := urlquery.NewClient(apikey)

		// Handle report data
		validActions := map[string]bool{
			"report":       true,
			"screenshot":   true,
			"domain_graph": true,
			"resource":     true,
		}
		if validActions[action] {
			// fmt.Printf("Fetching %s for Report: %s\n", action, report_id)

			// Fetch Report
			if action == "report" {
				report, err := client.GetReport(report_id)
				if err != nil {
					fmt.Println("Failed", err)
					os.Exit(1)
				}

				summary := viper.GetBool("summary")
				if summary {
					var parsed api.Report
					if err := json.Unmarshal(report.Bytes(), &parsed); err != nil {
						fmt.Println("Error parsing report:", err)
						os.Exit(1)
					}

					tmp := SummarizeReport(&parsed)
					fmt.Println(tmp)

				} else {
					reportFilename := fmt.Sprintf("report_%s.json", report_id)
					filePath := filepath.Join(output_directory, reportFilename)

					f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
					if err != nil {
						fmt.Println("Failed to write file:", err)
						os.Exit(1)
					}
					defer f.Close()
					f.Write(report.Bytes())
				}

			}

			// Fetch Domain Graph
			if action == "domain_graph" {
				domain_graph_filename := fmt.Sprintf("domain_graph_%s.gif", report_id)

				data_domain_graph, _ := client.GetDomainGraph(report_id)

				f, _ := os.OpenFile(output_directory+domain_graph_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

				resource_filename := fmt.Sprintf("resource_%s", hash)

				data_resource, err := client.GetResource(report_id, hash)
				if err != nil {
					fmt.Print("Error downloading resource:", err)
					return
				}
				fmt.Println("bytes:", len(data_resource))

				f, _ := os.OpenFile(output_directory+resource_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
				f.Write(data_resource)
				f.Close()

				return
			}

		}

	},
}
