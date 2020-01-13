package providers

import (
	"net/url"

	"golang.org/x/oauth2"
)

type oauthProvidersMap map[string]OAuthProvider

func (pm *oauthProvidersMap) register(p OAuthProvider) {
	OAuthProviders[p.Name] = p
}

var OAuthProviders = make(oauthProvidersMap)

type OAuthProvider struct {
	oauth2.Config
	Name        string
	UserDataURL *url.URL
}
