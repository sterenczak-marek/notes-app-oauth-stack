package internal

import (
	"html/template"
	"log"
	"net/http"
)

func ResponseHTML(rw http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(filename))
	if err := tmpl.Execute(rw, data); err != nil {
		log.Panicf("Unable to parse template file=%s. Error: %s", filename, err)
	}
}
