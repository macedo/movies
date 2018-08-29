package api

import (
	"encoding/json"
	"net/http"
)

// respondWithError uses respondWithJson in order to return error messages
// as JSON in a given http.ResponseWriter
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

// respondWithJson writes on w a JSON payload and also set the given http
// status code on it
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
