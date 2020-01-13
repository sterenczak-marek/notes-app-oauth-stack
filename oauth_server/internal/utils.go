package internal

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func responseHTML(rw http.ResponseWriter, filename string, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles(filename))
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
