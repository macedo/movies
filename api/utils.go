package api

import (
	"encoding/json"
	"net/http"
)

// respondWithError uses respondWithJSON in order to return error messages
// as JSON in a given http.ResponseWriter
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON writes on w a JSON payload and also set the given http
// status code on it
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
