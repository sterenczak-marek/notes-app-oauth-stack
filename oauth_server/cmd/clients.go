package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/cmd/cli"
)

var clientsHandler = &cli.ClientHandler{}
var clientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "Command for OAuth clients management",
}

var clientsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new OAuth client",
	Long: `This command creates an OAuth Client.
Example:
  oauth_server clients create --id 777 --secret 999 --redirectURI http://test.test
`,
	Run: clientsHandler.CreateClient,
}

func init() {
	rootCmd.AddCommand(clientsCmd)

	clientsCmd.AddCommand(clientsCreateCmd)

	clientsCmd.PersistentFlags().StringVar(&cli.ClientID, "id", "", "OAuth client ID")
	clientsCmd.MarkPersistentFlagRequired("id")

	clientsCreateCmd.Flags().StringVar(&cli.ClientSecret, "secret", "", "OAuth client secret")
	clientsCreateCmd.Flags().StringVar(&cli.RedirectURI, "redirectURI", "", "OAuth client redirect URI")
	clientsCreateCmd.MarkFlagRequired("redirectURI")
}
