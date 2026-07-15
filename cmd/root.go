package cmd

import (
	"os"

	"dynamic_runner_subsystem/internal/config"
	"dynamic_runner_subsystem/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dynamic_runner_subsystem",
	Short: "Manages dynamic Forgejo runners",
	Long:  `A subsystem for dynamic Forgejo runners and webhook provisioning.`,
}

// Execute runs the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize dependencies required by every command
	cobra.OnInitialize(func() {
		logger.Init()
		config.Load()
	})
}
