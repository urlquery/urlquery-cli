package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiKey string
var cfgFile string

func init() {
	// cobra.OnInitialize(initConfig)

	// Define API key flag but don't enforce it as required here
	rootCmd.PersistentFlags().String("apikey", "", "API Key (can be set via config file or flag)")
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "output", "", "Location to store downloaded fetched data (reports, screenshot, files)")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	// Config
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configUnsetCmd)

	// Submit
	submitCmd.Flags().String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:134.0) Gecko/20100101 Firefox/134.0", "Override default user agent")
	submitCmd.Flags().String("access", "public", "Override default access (public, restricted, private)")
	viper.BindPFlag("useragent", submitCmd.Flags().Lookup("useragent"))
	viper.BindPFlag("access", submitCmd.Flags().Lookup("access"))
	submitCmd.AddCommand(submitStatusCmd)

	// Search
	searchCmd.Flags().IntVar(&limitSearch, "limit", 10, "Maxium number of results to return")
	searchCmd.Flags().IntVar(&offsetSearch, "offset", 0, "Search  offset")
	// searchCmd.Flags().StringVar(&report_id, "type", "", "Search query (required)")

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(reputationCmd)
	rootCmd.AddCommand(searchCmd)
}

var rootCmd = &cobra.Command{
	Use:   "urlquery-cli",
	Short: "CLI for interacting with urlquery.net",
	Long:  `A command-line interface for querying and analyzing URLs via the urlquery API.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error finding home directory:", err)
				os.Exit(1)
			}
			viper.AddConfigPath(home)
			viper.SetConfigName(".urlquery-cli")
			viper.SetConfigType("yaml")
		}

		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err == nil {
			// fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

		// Override with flags **only if they are set**
		if cmd.Flags().Changed("apikey") {
			viper.Set("apikey", cmd.Flag("apikey").Value.String())
		}
		if cmd.Flags().Changed("output") {
			viper.Set("output", cmd.Flag("output").Value.String())
		}

		if cmd.Flags().Changed("base_url") {
			viper.Set("base_url", cmd.Flag("base_url").Value.String())
		}

		// Retrieve final API key value
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key is required. Set it via --apikey flag or in the config file.")
			os.Exit(1)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	var cfgFile string
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".urlquery" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".urlquery")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
