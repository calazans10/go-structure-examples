package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Returns the movies.")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Returns a movie.")
}

func getMovieReviews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Returns all reviews for a movie.")
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Adds a new movie.")
}

func addMovieReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Adds a new review for a movie.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}/reviews", getMovieReviews).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}/reviews", addMovieReview).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
