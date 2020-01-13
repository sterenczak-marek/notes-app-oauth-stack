package internal

import (
	"encoding/json"
	"net/http"
)

func ResponseError(rw http.ResponseWriter, errMsg string, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	_ = json.NewEncoder(rw).Encode(map[string]string{"error": errMsg})
}

func ResponseJSON(rw http.ResponseWriter, data interface{}, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	_ = json.NewEncoder(rw).Encode(data)
}
