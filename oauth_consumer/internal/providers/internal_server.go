package providers

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/oauth2"
)

func init() {
	v := os.Getenv("OAUTH_INTERNAL_SERVER_USER_DATA_URL")
	if len(v) < 1 {
		log.Fatalf("provide `OAUTH_INTERNAL_SERVER_USER_DATA_URL` env variable")
	}
	userDataURL, err := url.Parse(v)
	if err != nil {
		log.Fatalf("unable to parse user data url: %s", err)
	}

	OAuthProviders.register(
		OAuthProvider{
			Name: "internal",
			Config: oauth2.Config{
				ClientID:     os.Getenv("OAUTH_INTERNAL_SERVER_CLIENT_ID"),
				ClientSecret: os.Getenv("OAUTH_INTERNAL_SERVER_CLIENT_SECRET"),
				Scopes:       []string{"all"},
				RedirectURL:  fmt.Sprintf("%s/oauth/callback/internal", os.Getenv("APP_HOST")),
				Endpoint: oauth2.Endpoint{
					AuthURL:  os.Getenv("OAUTH_INTERNAL_SERVER_AUTH_URL"),
					TokenURL: os.Getenv("OAUTH_INTERNAL_SERVER_TOKEN_URL"),
				},
			},
			UserDataURL: userDataURL,
		},
	)
}
