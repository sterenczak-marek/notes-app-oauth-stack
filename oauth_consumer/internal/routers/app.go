package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/handlers"
)

func SetAppRoutes(router *mux.Router) *mux.Router {

	router.Handle(
		"/",
		negroni.New(
			negroni.HandlerFunc(internal.IsAuthenticated),
			negroni.Wrap(http.HandlerFunc(handlers.IndexPageHandler)),
		),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/login",
		handlers.LoginPageHandler,
	).Methods(http.MethodGet)

	router.Handle(
		"/logout",
		negroni.New(
			negroni.HandlerFunc(internal.IsAuthenticated),
			negroni.Wrap(http.HandlerFunc(handlers.LogoutPageHandler)),
		)).Methods("GET")

	return router
}
