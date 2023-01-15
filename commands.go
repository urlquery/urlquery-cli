package main

import (
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit [URL]",
	Short: "Submit a URL",
	Args:  cobra.ExactArgs(1),
	Run:   SubmitCmd,
}

var reportCmd = &cobra.Command{
	Use:   "report [report_id]",
	Short: "Retrive a report",
	Args:  cobra.ExactArgs(1),
	Run:   ReportCmd,
}

var screenshotCmd = &cobra.Command{
	Use:   "screenshot [report_id]",
	Short: "Download screenshot",
	Args:  cobra.ExactArgs(1),
	Run:   ScreenshotCmd,
}

var repCheckCmd = &cobra.Command{
	Use:   "rep-check [url/ip]",
	Short: "Check the reputation of a URL or IP",
	Args:  cobra.ExactArgs(1),
	Run:   RepCheckCmd,
}
