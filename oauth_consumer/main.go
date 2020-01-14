package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/routers"
)

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}
	log.Printf("Listening for connections on port: %s\n", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, n))
}
