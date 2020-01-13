package internal

import (
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"gopkg.in/oauth2.v3/server"
)

func ValidateClientCredentialsMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	_, _, err := server.ClientBasicHandler(req)
	if err != nil {
		data, status, _ := config.OAuthServer.GetErrorData(err)
		responseJSON(rw, data, status)
		return
	}

	next(rw, req)
}
