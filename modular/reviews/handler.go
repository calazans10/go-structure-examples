package reviews

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetMovieReviews returns all reviews for a movie
func GetMovieReviews(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading reviews"
		vars := mux.Vars(r)
		id := vars["id"]

		data, err := service.FindMovieReviews(id)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// AddMovieReview adds a new review for a movie
func AddMovieReview(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding a review for a movie"
		vars := mux.Vars(r)
		id := vars["id"]

		review := Review{
			MovieID: id,
		}

		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := service.SaveMovieReview(&review)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithEmpty(w, http.StatusCreated)
	})
}

func respondWithJSON(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithEmpty(w http.ResponseWriter, code int) {
	respondWithJSON(w, "", code)
}
