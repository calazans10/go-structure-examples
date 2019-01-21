package movies

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetMovies returns the movies
func GetMovies(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading movies"

		data, err := service.FindMovies()
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// GetMovie returns a movie
func GetMovie(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading a movie"
		vars := mux.Vars(r)
		id := vars["id"]

		data, err := service.FindMovie(id)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// AddMovie adds a new movie
func AddMovie(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding a movie"
		movie := Movie{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := service.SaveMovie(&movie)
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
