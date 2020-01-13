package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"gopkg.in/oauth2.v3/models"
)

type ClientHandler struct{}

var ClientID = ""
var ClientSecret = ""
var RedirectURI = ""

func (h *ClientHandler) CreateClient(_ *cobra.Command, _ []string) {
	client, err := config.OAuthManager.GetClient(ClientID)
	if err == nil && client != nil {
		log.Fatalf("Client with ID=%s already exists\n", ClientID)
	}

	var echoSecret bool
	if ClientSecret == "" {
		ClientSecret = generateSecret(26)
		echoSecret = true
	}

	err = config.OAuthClientStore.Create(&models.Client{
		ID:     ClientID,
		Secret: ClientSecret,
		Domain: RedirectURI,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
	}
	client, err = config.OAuthManager.GetClient(ClientID)
	if err != nil {
		log.Fatalf("Failed to fetch created client: %s", err)
	}
	fmt.Printf("Created OAuth Client ID: %s\n", client.GetID())
	if echoSecret {
		log.Fatalf("OAuth Client Secret: %s\n", ClientSecret)
	}
}
