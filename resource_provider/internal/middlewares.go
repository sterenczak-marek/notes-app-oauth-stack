package internal

import (
	"log"
	"net/http"
	"strings"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal/validators"
)

const bearerPrefix = "bearer "

func ValidateAccessTokenMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	header := req.Header.Get("Authorization")
	n := len(bearerPrefix)
	if len(header) < n || strings.ToLower(header[:n]) != bearerPrefix {
		http.Error(rw, "Authorization header invalid. Valid format is `Authorization: bearer <token>`", http.StatusUnauthorized)
		return
	}
	token := header[n:]

	providerName := req.Header.Get("X-Provider")

	validator, ok := validators.OAuthValidators[providerName]
	if !ok {
		log.Printf("Unable to get validator for provider name=%s", providerName)
		ResponseError(rw, "Invalid access token", http.StatusUnauthorized)
		return
	}

	if validator(token) {
		next(rw, req)
	} else {
		ResponseError(rw, "Invalid access token", http.StatusUnauthorized)
	}
}
