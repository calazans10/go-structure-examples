package main

import (
	"fmt"
	"net/http"
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
