package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal"
	"github.com/urfave/negroni"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/handlers"
)

func SetProxyRoutes(router *mux.Router) *mux.Router {
	middlewares := negroni.New(
		negroni.HandlerFunc(internal.IsAuthenticated),
	)

	router.Handle(
		"/notes",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.ProxyToResourceProvider)),
		),
	).Methods(http.MethodGet, http.MethodPost)

	router.Handle(
		"/notes/{noteUUID}",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.ProxyToResourceProvider)),
		),
	).Methods(http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete)

	return router
}
