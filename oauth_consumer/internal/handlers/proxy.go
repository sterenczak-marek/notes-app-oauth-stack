package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/config"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal"
)

func ProxyToResourceProvider(rw http.ResponseWriter, req *http.Request) {
	user := internal.MustGetUserFromContext(req.Context())
	resourceProviderURL := config.ResourceProviderURL
	proxy := httputil.NewSingleHostReverseProxy(resourceProviderURL)

	req.URL.Host = resourceProviderURL.Host
	req.URL.Scheme = resourceProviderURL.Scheme
	req.Host = resourceProviderURL.Host
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("X-User", user.Email)
	req.Header.Set("X-Provider", user.ProviderName)

	accessToken := fmt.Sprintf("Bearer %s", user.AccessToken)
	req.Header.Set("Authorization", accessToken)

	proxy.ServeHTTP(rw, req)
}
