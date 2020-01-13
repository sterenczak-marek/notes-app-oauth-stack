package handlers

import (
	"log"
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/config"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal"
)

func IndexPageHandler(rw http.ResponseWriter, req *http.Request) {
	user := internal.MustGetUserFromContext(req.Context())

	internal.ResponseHTML(rw, "templates/index.html", user)
}

func LoginPageHandler(rw http.ResponseWriter, _ *http.Request) {
	internal.ResponseHTML(rw, "templates/login.html", nil)
}

func LogoutPageHandler(rw http.ResponseWriter, req *http.Request) {
	store, err := config.SessionStore.Get(req, config.CookieSessionKey)
	if err != nil {
		log.Panicf("unable to get session data: %s", err)
	}

	store.Options.MaxAge = -1
	if err = store.Save(req, rw); err != nil {
		log.Printf("Unable to delete session data: %s", err)
		return
	}
	http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
}
