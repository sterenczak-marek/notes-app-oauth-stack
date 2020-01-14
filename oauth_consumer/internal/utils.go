package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/config"
)

func ResponseHTML(rw http.ResponseWriter, filename string, data interface{}) {
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
