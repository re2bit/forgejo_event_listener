package cmd

import (
	"dynamic_runner_subsystem/internal/api"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		// Starts the Gin server
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
