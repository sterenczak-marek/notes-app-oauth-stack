package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetAppRoutes(router)
	router = SetOAuthRoutes(router)
	router = SetProxyRoutes(router)

	return router
}
