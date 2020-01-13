package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal/routers"
)

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	log.Println("Listening for connections on port: 80")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", n))
}
