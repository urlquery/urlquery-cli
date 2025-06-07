package output

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// FormatJSON formats and prints JSON with proper indentation
func FormatJSON(data interface{}) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(output))
	return nil
}

// PrintSuccess prints a success message in green
func PrintSuccess(message string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("✅ %s\n", green(message))
}

// PrintError prints an error message in red
func PrintError(message string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("❌ %s\n", red(message))
}

// PrintWarning prints a warning message in yellow
func PrintWarning(message string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("⚠️  %s\n", yellow(message))
}

// PrintInfo prints an info message in blue
func PrintInfo(message string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("ℹ️  %s\n", blue(message))
}

// PrintHeader prints a formatted header
func PrintHeader(title string) {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Printf("\n%s\n", bold(title))
	fmt.Println(strings.Repeat("─", len(title)))
}

// PrintKeyValue prints a key-value pair with consistent formatting
func PrintKeyValue(key, value string) {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Printf("%s: %s\n", bold(key), value)
}

// PrintTable prints data in a simple table format
func PrintTable(headers []string, rows [][]string) {
	if len(headers) == 0 || len(rows) == 0 {
		return
	}

	// Calculate column widths
	widths := make([]int, len(headers))
	for i, header := range headers {
		widths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Print header
	bold := color.New(color.Bold).SprintFunc()
	for i, header := range headers {
		fmt.Printf("%-*s", widths[i]+2, bold(header))
	}
	fmt.Println()

	// Print separator
	for _, width := range widths {
		fmt.Print(strings.Repeat("─", width+2))
	}
	fmt.Println()

	// Print rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				fmt.Printf("%-*s", widths[i]+2, cell)
			}
		}
		fmt.Println()
	}
}
