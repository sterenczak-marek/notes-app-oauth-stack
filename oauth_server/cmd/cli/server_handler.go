package cli

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/internal"
	"github.com/urfave/negroni"
)

type ServerHandler struct{}

func (h *ServerHandler) HandleRequests(_ *cobra.Command, _ []string) {
	router := internal.NewRouter()

	n := negroni.Classic()
	n.UseHandler(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}

	log.Printf("Server is running at %s port.\n", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, n))
}

type JSONPanicFormatter struct{}

func (t *JSONPanicFormatter) FormatPanicError(rw http.ResponseWriter, _ *http.Request, _ *negroni.PanicInformation) {
	rw.Header().Set("Content-type", "application/json; charset=UTF-8")

	jsonData, err := json.Marshal(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Internal server error",
	})
	if err != nil {
		// cannot raise panic one more time
		log.Printf("Unable to encode data to JSON. Reason: %s", err)
	}
	_, err = rw.Write(jsonData)
	if err != nil {
		log.Printf("Unable to write data to response, reason: %s", err)
	}
}
