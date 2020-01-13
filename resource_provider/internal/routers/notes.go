package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal"
	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal/handlers"
)

func SetNotesRoutes(router *mux.Router) *mux.Router {

	middlewares := negroni.New(
		negroni.HandlerFunc(internal.ValidateAccessTokenMiddleware),
	)

	router.Handle(
		"/notes",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.NoteListHandler)),
		)).Methods(http.MethodGet)

	router.Handle(
		"/notes",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.NoteCreateHandler)),
		)).Methods(http.MethodPost)

	router.Handle(
		"/notes/{noteUUID}",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.NoteDetailHandler)),
		)).Methods(http.MethodGet)

	router.Handle(
		"/notes/{noteUUID}",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.NoteUpdateHandler)),
		)).Methods(http.MethodPatch, http.MethodPut)

	router.Handle(
		"/notes/{noteUUID}",
		middlewares.With(
			negroni.Wrap(http.HandlerFunc(handlers.NoteDeleteHandler)),
		)).Methods(http.MethodDelete)

	return router
}
