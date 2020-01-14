package internal

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
)

func responseHTML(rw http.ResponseWriter, filename string, data map[string]interface{}) {
	content, err := config.HTMLTemplateBox.FindString(filename)
	if err != nil {
		log.Panic(err)
	}
	tmpl := template.Must(
		template.New(filename).Parse(content),
	)
	if err := tmpl.Execute(rw, data); err != nil {
		log.Panicf("Unable to parse template file=%s. Error: %s", filename, err)
	}
}

func responseError(rw http.ResponseWriter, errMsg string, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	_ = json.NewEncoder(rw).Encode(map[string]string{"error": errMsg})
}

func responseJSON(rw http.ResponseWriter, data interface{}, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	_ = json.NewEncoder(rw).Encode(data)
}
