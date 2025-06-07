package main

import (
	"github.com/urlquery/urlquery-cli/cmd"
)

// Build information (set via ldflags during build)
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func main() {
	// Set version information for the CLI
	cmd.SetVersionInfo(Version, BuildTime, GitCommit)
	cmd.Execute()
}
