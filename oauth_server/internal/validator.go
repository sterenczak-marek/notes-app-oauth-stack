package internal

import (
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"gopkg.in/oauth2.v3/errors"
)

// exactTheSameURIValidator validates that redirectURI is exactly the same as baseURI
func exactTheSameURIValidator(baseURI string, redirectURI string) error {
	if baseURI != redirectURI {
		return errors.ErrInvalidRedirectURI
	}
	return nil
}

func init() {
	// override redirect URI validator - has to be exactly the same URIs, not only suffix the same
	config.OAuthManager.SetValidateURIHandler(exactTheSameURIValidator)
}
