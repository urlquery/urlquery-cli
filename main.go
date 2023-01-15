package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"urlquery-cli/api/public"
	"urlquery-cli/types"

	"github.com/apoorvam/goterminal"
	"github.com/spf13/cobra"
)

var apikey string
var output string

func main() {
	var rootCmd = &cobra.Command{Use: "urlquery-cli", DisableAutoGenTag: true}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&apikey, "apikey", "", "API key")
	rootCmd.PersistentFlags().StringVar(&output, "output", "stdout", "Output (stdout, file)")

	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(screenshotCmd)
	rootCmd.AddCommand(repCheckCmd)

	rootCmd.Execute()
}

func SubmitCmd(cmd *cobra.Command, args []string) {
	data, err := public.NewDefaultRequest().SubmitURL(args[0])
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	var q types.Queue
	json.Unmarshal([]byte(data), &q)
	qid := q.QueueID
	fmt.Printf("URL submitted!\n")

	writer := goterminal.New(os.Stdout)
	for q.Status != "done" {
		fmt.Fprintf(writer, "Current status: %s\n", q.Status)
		writer.Print()

		time.Sleep(time.Millisecond * 1000)
		output, _ := public.NewDefaultRequest().GetQueueStatus(qid)
		json.Unmarshal([]byte(output), &q)
		writer.Clear()
	}
	writer.Reset()

	fmt.Printf("Report finished:\n")
	fmt.Printf("  - %s\n", q.ReportID)
	fmt.Printf("  - https://urlquery.net/report/%s\n", q.ReportID)

	fmt.Printf("\nTo grab the report run:\n")
	fmt.Printf("   urlquery-cli report %s\n", q.ReportID)
}

func ReportCmd(cmd *cobra.Command, args []string) {
	report_id := args[0]
	data, err := public.NewDefaultRequest().GetReport(report_id)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	if output == "file" {
		filename := fmt.Sprintf("urlquery_report_%s.json", report_id)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		fmt.Printf("File written to: %s\n", filename)
		file.WriteString(data)
	} else {
		var report types.ReportOverview
		json.Unmarshal([]byte(data), &report)

		fmt.Printf("\n\n")
		fmt.Printf("Report URL: https://urlquery.net/report/%s\n", report.ID)
		fmt.Printf(" --- \n")
		fmt.Printf("URL:  %s\n", strings.ReplaceAll(report.Url.Addr, ".", "[.]"))
		fmt.Printf("IP:   %s (%s)\n", report.Ip.Addr, report.Ip.Country)
		fmt.Printf("Tags: %s\n", strings.Join(report.Tags, ", "))
		fmt.Printf("Alerts\n")
		fmt.Printf("  IDS:       %d\n", report.IDSAlertCount)
		fmt.Printf("  urlquery:  %d\n", report.UrlQueryAlertCount)
		fmt.Printf("  blocklist: %d\n", report.BlocklistAlertCount)
	}

}

func ScreenshotCmd(cmd *cobra.Command, args []string) {
	report_id := args[0]
	data, err := public.NewDefaultRequest().GetScreenshot(report_id)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	filename := fmt.Sprintf("urlquery_screenshot_%s.png", report_id)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	file.Write(data)
	fmt.Printf("Screenshot written to: %s\n", filename)
}

func RepCheckCmd(cmd *cobra.Command, args []string) {
	data, _ := public.NewDefaultRequest().CheckReputation(args[0])
	fmt.Println(data)
}
