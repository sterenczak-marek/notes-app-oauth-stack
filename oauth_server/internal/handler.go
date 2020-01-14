package internal

import (
	"log"
	"net/http"
	"net/url"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/internal/models"
	"gopkg.in/oauth2.v3"
)

const (
	AccessTokenHint  = "access_token"
	RefreshTokenHint = "refresh_token"
	TokenType        = "Bearer"

	LoggedUsersKey        = "loggedUserUUID"
	AuthenticatedUsersKey = "authenticatedUserUUID"
)

func init() {
	config.OAuthServer.SetUserAuthorizationHandler(userAuthorizeHandler)
}

func userAuthorizeHandler(rw http.ResponseWriter, req *http.Request) (_ string, err error) {
	store, _ := config.SessionStore.Get(req, config.SessionCookieKey)
	userUUID, ok := store.Values[AuthenticatedUsersKey]
	if !ok {
		if err = req.ParseForm(); err != nil {
			log.Printf("Unable to parse request form: %s", err)
			return
		}
		store.Values["ReturnFormValues"] = req.Form
		if err = store.Save(req, rw); err != nil {
			log.Printf("Unable to save session data: %s", err)
			return
		}

		rw.Header().Set("Location", "/login")
		rw.WriteHeader(http.StatusFound)
		return
	}
	return userUUID.(string), nil
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	if err := req.ParseForm(); err != nil {
		log.Panicf("Unable to parse request form: %s", err)
	}

	store, err := config.SessionStore.Get(req, config.SessionCookieKey)
	if err != nil {
		log.Panic(err)
	}
	if _, ok := store.Values["ReturnFormValues"]; !ok {
		http.Error(rw, "Only Oauth authorization code is allowed", http.StatusNotFound)
		return
	}
	if req.Method == "POST" {
		user, err := models.Authenticate(req.FormValue("email"), req.FormValue("password"))
		if err == nil {
			store.Values[LoggedUsersKey] = user.UUID.String()
			if err = store.Save(req, rw); err != nil {
				log.Panicf("Error saving session: %s", err)
			}
			http.Redirect(rw, req, "/auth", http.StatusFound)
			return
		}
		data["errors"] = []string{err.Error()}
	}
	responseHTML(rw, "login.html", data)
}

func authHandler(rw http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})

	store, err := config.SessionStore.Get(req, config.SessionCookieKey)
	if err != nil {
		log.Panic(err)
	}
	if _, ok := store.Values[LoggedUsersKey]; !ok {
		http.Redirect(rw, req, "/login", http.StatusFound)
		return
	}
	if v, ok := store.Values["ReturnFormValues"]; ok {
		redirectURI := v.(url.Values)["redirect_uri"][0]
		if redirectApp, err := url.Parse(redirectURI); err == nil {
			data["RedirectDomain"] = redirectApp.Host
		}
	}

	if req.Method == http.MethodPost {
		var form url.Values
		if v, ok := store.Values["ReturnFormValues"]; ok {
			form = v.(url.Values)
		}
		u := new(url.URL)
		u.Path = "/authorize"
		u.RawQuery = form.Encode()

		http.Redirect(rw, req, u.String(), http.StatusFound)
		delete(store.Values, "ReturnFormValues")

		if v, ok := store.Values[LoggedUsersKey]; ok {
			store.Values[AuthenticatedUsersKey] = v
		}
		if err = store.Save(req, rw); err != nil {
			log.Panicf("Error saving session: %s", err)
		}
		return
	}
	responseHTML(rw, "auth.html", data)
}

func validateTokenHandler(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Panicf("unable to parse form data: %s", err)
	}

	token := req.Form.Get("token")
	if token == "" {
		responseError(rw, "Token missing", http.StatusBadRequest)
		return
	}

	tokenTypeHint := req.Form.Get("token_type_hint")
	if tokenTypeHint == "" {
		tokenTypeHint = AccessTokenHint
	}

	var getTokenFunc func(string) (oauth2.TokenInfo, error)
	var getExpiresAt func(oauth2.TokenInfo) int

	switch tokenTypeHint {
	case AccessTokenHint:
		getTokenFunc = config.OAuthServer.Manager.LoadAccessToken
		getExpiresAt = func(tokenInfo oauth2.TokenInfo) int {
			return int(tokenInfo.GetAccessCreateAt().Add(tokenInfo.GetAccessExpiresIn()).Unix())
		}
	case RefreshTokenHint:
		getTokenFunc = config.OAuthServer.Manager.LoadRefreshToken
		getExpiresAt = func(tokenInfo oauth2.TokenInfo) int {
			return int(tokenInfo.GetRefreshCreateAt().Add(tokenInfo.GetRefreshExpiresIn()).Unix())
		}
	default:
		responseError(rw, "Token hint invalid", http.StatusBadRequest)
		return
	}

	tokenInfo, _ := getTokenFunc(token)
	var resp IntrospectResponse
	if tokenInfo == nil {
		resp = IntrospectResponse{Active: false}
	} else {
		resp = IntrospectResponse{
			Active:    true,
			Scope:     tokenInfo.GetScope(),
			ClientID:  tokenInfo.GetClientID(),
			TokenType: TokenType,
			ExpiresAt: getExpiresAt(tokenInfo),
		}
	}
	responseJSON(rw, resp, http.StatusOK)
}

func userUserDataHandler(rw http.ResponseWriter, req *http.Request) {
	token, err := config.OAuthServer.ValidationBearerToken(req)
	if err != nil {
		responseError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.GetUserByID(token.GetUserID())
	responseJSON(rw, user, http.StatusOK)
}

type IntrospectResponse struct {
	Active    bool   `json:"active"`
	Scope     string `json:"scope,omitempty"`
	ClientID  string `json:"client_id,omitempty"`
	Username  string `json:"username,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	ExpiresAt int    `json:"exp,omitempty"`
}
