package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler defines the handler structure
type Handler struct {
	service *Service
}

// NewHandler creates a new handler
func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

// GetMovies returns the movies
func (h *Handler) GetMovies() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading movies"

		data, err := h.service.FindMovies()
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// GetMovie returns a movie
func (h *Handler) GetMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading a movie"
		vars := mux.Vars(r)
		id := vars["id"]

		data, err := h.service.FindMovie(id)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// GetMovieReviews returns all reviews for a movie
func (h *Handler) GetMovieReviews() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading reviews"
		vars := mux.Vars(r)
		id := vars["id"]

		data, err := h.service.FindMovieReviews(id)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, data, http.StatusOK)
	})
}

// AddMovie adds a new movie
func (h *Handler) AddMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding a movie"
		movie := Movie{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := h.service.SaveMovie(&movie)
		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		respondWithEmpty(w, http.StatusCreated)
	})
}

// AddMovieReview adds a new review for a movie
func (h *Handler) AddMovieReview() http.Handler {
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

		err := h.service.SaveMovieReview(&review)
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
