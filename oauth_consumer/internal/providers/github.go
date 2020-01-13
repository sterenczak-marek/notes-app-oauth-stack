package providers

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func init() {
	v := os.Getenv("GITHUB_USER_DATA_URL")
	if len(v) < 1 {
		log.Fatalf("provide `GITHUB_USER_DATA_URL` env variable")
	}
	userDataURL, err := url.Parse(v)
	if err != nil {
		log.Fatalf("unable to parse user data url: %s", err)
	}

	OAuthProviders.register(
		OAuthProvider{
			Name: "github",
			Config: oauth2.Config{
				ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
				Scopes:       []string{"all"},
				RedirectURL:  fmt.Sprintf("%s/oauth/callback/github", os.Getenv("APP_HOST")),
				Endpoint:     github.Endpoint,
			},
			UserDataURL: userDataURL,
		},
	)
}
