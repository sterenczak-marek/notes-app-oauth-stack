package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/config"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/models"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/providers"
)

func OAuthRedirect(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	provider, err := getProviderConfig(vars["provider"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	state := getRandomState()
	store, err := config.SessionStore.Get(req, config.CookieSessionKey)
	if err != nil {
		log.Panic(err)
	}
	store.Values["state"] = state
	if err := store.Save(req, rw); err != nil {
		log.Panic(err)
	}
	u := provider.AuthCodeURL(state)
	http.Redirect(rw, req, u, http.StatusTemporaryRedirect)
}

func OAuthCallback(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Panic(err)
	}
	vars := mux.Vars(req)
	provider, err := getProviderConfig(vars["provider"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	store, err := config.SessionStore.Get(req, config.CookieSessionKey)
	if err != nil {
		log.Panic(err)
	}
	state, ok := store.Values["state"]
	if !ok || req.FormValue("state") != state {
		http.Error(rw, "Invalid state", http.StatusBadRequest)
		return
	}
	code := req.FormValue("code")
	if code == "" {
		http.Error(rw, "Code parameter is required", http.StatusBadRequest)
		return
	}

	token, err := provider.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Code exchange with oauth server provider failed: %s", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := provider.Client(context.Background(), token)
	resp, err := client.Get(provider.UserDataURL.String())
	if err != nil {
		log.Printf("Getting user data information failed: %s", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	var content map[string]interface{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		log.Panic(err)
	}
	delete(store.Values, "state")

	if resp.StatusCode == http.StatusOK {
		store.Values["user"] = &models.User{
			Email:        content["email"].(string),
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			ProviderName: provider.Name,
		}
		if err := store.Save(req, rw); err != nil {
			log.Panic(err)
		}
		http.Redirect(rw, req, "/", http.StatusFound)
	} else {
		log.Printf("internal server error: %s", content["error_description"])
		log.Panic(fmt.Errorf(content["error"].(string)))
	}
}

func getProviderConfig(providerName string) (*providers.OAuthProvider, error) {
	provider, ok := providers.OAuthProviders[providerName]
	if !ok {
		return nil, fmt.Errorf("unknown oauth provider: %s", providerName)
	}
	return &provider, nil
}

func getRandomState() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicf("Unable to generate random state: %s", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
