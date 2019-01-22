package helpers

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON writes a json response format
func RespondWithJSON(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithEmpty writes a json response format with an empty body
func RespondWithEmpty(w http.ResponseWriter, code int) {
	RespondWithJSON(w, "", code)
}
