package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or edit CLI config (e.g., API key, default output)",
	Long: `Manage default settings used by urlquery-cli

Configuration values are stored in a YAML file located at ~/.urlquery-cli.yaml
This includes settings like the API key and default output directory.

Use 'urlquery-cli config show' to view current settings, and 'urlquery-cli config set <key> <value>' to update them.

Available keys:
  - apikey       Your urlquery API key
  - output       Default directory to save downloaded reports or files
  - useragent    Default useragent to use for submissions
  - access       Set default access for submitted URL (public, restricted, private)

Examples:
  urlquery-cli config show
  urlquery-cli config set apikey abc123
  urlquery-cli config set output ./downloads
  urlquery-cli config set access public
`,
}

// Show command (config show)
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Configuration:")
		for _, key := range viper.AllKeys() {
			fmt.Printf("  %s: %v\n", key, viper.Get(key))
		}
	},
}

var allowedConfigKeys = map[string]bool{
	"apikey": true,
	"output": true,
	"access": true,
}

var allowedAccessValues = map[string]bool{
	"public":     true,
	"restricted": true,
	"private":    true,
}

// Set command (config set <key> <value>)
var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long: `Set a configuration value for urlquery-cli.

Configuration settings are saved in ~/.urlquery-cli.yaml and are used as defaults
for CLI operations like URL submissions and result downloads.

Available keys:
  - apikey       Your urlquery API key
  - output       Default directory to save downloaded reports or files
  - useragent    Default User-Agent string for URL submissions
  - access       Default visibility for submitted URLs: public, restricted, or private

Examples:
  urlquery-cli config set apikey abc123
  urlquery-cli config set output ./downloads
  urlquery-cli config set useragent "curl/7.81.0"
  urlquery-cli config set access restricted`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		if !allowedConfigKeys[key] {
			fmt.Printf("Error: unsupported config key '%s'\n", key)
			os.Exit(1)
		}

		if key == "access" && !allowedAccessValues[value] {
			fmt.Printf("Error: invalid value for 'access'. Must be one of: public, restricted, private\n")
			os.Exit(1)
		}

		configFile := viper.ConfigFileUsed()
		if configFile == "" {
			home, _ := os.UserHomeDir()
			configFile = home + "/.urlquery-cli.yaml"
			viper.SetConfigFile(configFile)
		}

		viper.Set(key, value)

		// Save changes
		if err := viper.WriteConfigAs(configFile); err != nil {
			fmt.Println("Error saving config:", err)
			os.Exit(1)
		}

		fmt.Printf("Config updated: %s = %s\n", key, value)
	},
}

var configUnsetCmd = &cobra.Command{
	Use:   "unset <key>",
	Short: "Remove a configuration value",
	Long: `Remove a configuration key from urlquery-cli settings.

This deletes the specified key from the config file (~/.urlquery-cli.yaml),
which will cause the CLI to fall back to default behavior for that setting.

You can use this to clear values like API keys, output directories, or access modes.

Examples:
	urlquery-cli config unset apikey
	urlquery-cli config unset useragent`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		configFile := viper.ConfigFileUsed()
		if configFile == "" {
			home, _ := os.UserHomeDir()
			configFile = home + "/.urlquery-cli.yaml"
			viper.SetConfigFile(configFile)
		}

		if !viper.IsSet(key) {
			fmt.Printf("Config key '%s' is not set.\n", key)
			return
		}

		viper.Set(key, nil)  // Clear in memory
		viper.ReadInConfig() // Load existing config
		allSettings := viper.AllSettings()
		delete(allSettings, key)

		// Write updated config without the key
		var buf bytes.Buffer
		enc := yaml.NewEncoder(&buf)
		defer enc.Close()
		if err := enc.Encode(allSettings); err != nil {
			fmt.Println("Error encoding config:", err)
			os.Exit(1)
		}

		if err := os.WriteFile(configFile, buf.Bytes(), 0644); err != nil {
			fmt.Println("Error writing config:", err)
			os.Exit(1)
		}

		fmt.Printf("Config key '%s' has been removed.\n", key)
	},
}
