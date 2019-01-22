package api

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

// CheckMovieID verifies if a movie id is a valid uuid
func CheckMovieID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, err := uuid.FromString(vars["id"])
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}
		h.ServeHTTP(w, r)

	})
}
