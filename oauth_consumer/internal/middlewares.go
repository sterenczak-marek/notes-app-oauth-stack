package internal

import (
	"log"
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/config"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/models"
)

func IsAuthenticated(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	session, err := config.SessionStore.Get(req, config.CookieSessionKey)
	if err != nil {
		log.Panicf("unable to get session data: %s", err)
	}

	if user, ok := session.Values["user"]; !ok {
		http.Redirect(rw, req, "/login", http.StatusFound)
	} else {
		ctx := NewUserContext(req.Context(), user.(*models.User))
		next(rw, req.WithContext(ctx))
	}
}
