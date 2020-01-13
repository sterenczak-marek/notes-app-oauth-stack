package internal

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"github.com/urfave/negroni"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/auth", authHandler)

	router.Handle(
		"/validate",
		negroni.New(
			negroni.HandlerFunc(ValidateClientCredentialsMiddleware),
			negroni.WrapFunc(validateTokenHandler),
		),
	)
	router.HandleFunc(
		"/user",
		userUserDataHandler,
	)

	router.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := config.OAuthServer.HandleAuthorizeRequest(w, r)
		if err != nil {
			log.Panicf("Error occured on OAuth Server proccessing `authorize` request: %s", err)
		}
	})

	router.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := config.OAuthServer.HandleTokenRequest(w, r)
		if err != nil {
			log.Panicf("Error occured on OAuth Server proccessing `token` request: %s", err)
		}
	})

	return router
}
