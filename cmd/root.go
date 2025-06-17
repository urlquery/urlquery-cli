package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var outputSummary bool

// Version information
var (
	version   = "dev"
	buildTime = "unknown"
	gitCommit = "unknown"
)

// SetVersionInfo sets the version information for the CLI
func SetVersionInfo(v, bt, gc string) {
	version = v
	buildTime = bt
	gitCommit = gc
	rootCmd.Version = v
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config file path flag (not used as a Viper key)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file (default is $HOME/.urlquery-cli.yaml)")

	// Global flags
	rootCmd.PersistentFlags().String("apikey", "", "API Key (can also be set via config file or URLQUERY_APIKEY env var)")
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))

	rootCmd.PersistentFlags().String("output", "", "Location to store downloaded data (reports, screenshots, files)")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().Bool("summary", false, "Show a summary output instead of full json")
	viper.BindPFlag("summary", rootCmd.PersistentFlags().Lookup("summary"))

	// env settings
	viper.SetEnvPrefix("urlquery")
	viper.AutomaticEnv()

	// Submit command flags
	submitCmd.Flags().String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:134.0) Gecko/20100101 Firefox/134.0", "Override default user agent")
	submitCmd.Flags().String("access", "public", "Override default access level (public, restricted, private)")
	viper.BindPFlag("useragent", submitCmd.Flags().Lookup("useragent"))
	viper.BindPFlag("access", submitCmd.Flags().Lookup("access"))
	submitCmd.AddCommand(submitStatusCmd)

	// Search command flags
	searchCmd.Flags().IntVar(&limitSearch, "limit", 10, "Maximum number of results to return")
	searchCmd.Flags().IntVar(&offsetSearch, "offset", 0, "Offset for paginated search results")

	reportCmd.Flags().BoolVar(&outputSummary, "summary", false, "Show summary output instead of full report")

	// Register commands
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(reputationCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(versionCmd)

	// Add subcommands
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configUnsetCmd)
}

var rootCmd = &cobra.Command{
	Use:   "urlquery-cli",
	Short: "CLI for interacting with urlquery.net",
	Long:  `A command-line interface for querying and analyzing URLs via the urlquery API.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()

		// Skip API key check for config-related commands
		if cmd.Name() == "config" || cmd.HasParent() && cmd.Parent().Name() == "config" {
			return
		}

		// Override with flags **only if they are set**
		if cmd.Flags().Changed("apikey") {
			viper.Set("apikey", cmd.Flag("apikey").Value.String())
		}

		if cmd.Flags().Changed("output") {
			viper.Set("output", cmd.Flag("output").Value.String())
		}

		if cmd.Flags().Changed("summary") {
			viper.Set("summary", cmd.Flag("summary").Value.String())
		}

		if cmd.Flags().Changed("apigw_base") {
			viper.Set("apigw_base", cmd.Flag("apigw_base").Value.String())
		}

		// Check API key value
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key is required. Set it via 'config set apikey <value>' or use the --apikey flag.")
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

	viper.SetEnvPrefix("urlquery")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
