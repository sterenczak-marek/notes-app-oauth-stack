package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/cmd/cli"
)

var usersHandler = &cli.UserHandler{}
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Command for OAuth server users management",
}

var usersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new OAuth server user",
	Run:   usersHandler.CreateUser,
}

func init() {
	rootCmd.AddCommand(usersCmd)

	usersCmd.AddCommand(usersCreateCmd)

	usersCreateCmd.Flags().StringVar(&cli.Email, "email", "", "OAuth server user email")
	usersCreateCmd.MarkFlagRequired("email")

	usersCreateCmd.Flags().StringVar(&cli.Password, "password", "", "OAuth server user password")

}
