package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/cmd/cli"
)

var serverHandler = &cli.ServerHandler{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oauth_server",
	Short: "OAuth server application",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: serverHandler.HandleRequests,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
