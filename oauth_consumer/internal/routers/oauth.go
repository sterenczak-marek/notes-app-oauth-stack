package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/handlers"
)

func SetOAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(
		"/oauth/redirect/{provider}",
		handlers.OAuthRedirect,
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/oauth/callback/{provider}",
		handlers.OAuthCallback,
	).Methods(http.MethodGet)

	return router
}
