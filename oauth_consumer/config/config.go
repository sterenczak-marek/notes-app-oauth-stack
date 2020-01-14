package config

import (
	"log"
	"net/url"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var (
	SessionStore        *sessions.CookieStore
	ResourceProviderURL *url.URL
	HTMLTemplateBox     = packr.New("htmlTemplates", "../templates")
)

const CookieSessionKey = "consumer-app-session"

func init() {
	initSessionStore()
	getResourceProvider()
}

func initSessionStore() {
	sessionSecret := os.Getenv("SESSION_KEY")
	if sessionSecret == "" {
		log.Fatalf("provide `SESSION_KEY` env variable")
	}
	store := sessions.NewCookieStore([]byte(sessionSecret))
	SessionStore = store
}

func getResourceProvider() {
	resourceProviderURL := os.Getenv("RESOURCE_PROVIDER_URL")
	if resourceProviderURL == "" {
		log.Fatalf("provide `RESOURCE_PROVIDER_URL` env variable")
	}

	rpURL, err := url.Parse(resourceProviderURL)
	if err != nil {
		log.Fatalf("Unable to parse resource provider URL")
	}
	ResourceProviderURL = rpURL
}
