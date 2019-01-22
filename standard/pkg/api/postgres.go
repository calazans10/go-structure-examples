package api

import (
	"github.com/jmoiron/sqlx"
)

type repo struct {
	DB *sqlx.DB
}

// NewPostgresRepository creates a new repository
func NewPostgresRepository(db *sqlx.DB) Repository {
	return &repo{
		DB: db,
	}
}

// FindMovie returns a movie
func (r *repo) FindMovie(id string) (*Movie, error) {
	result := Movie{}
	err := r.DB.Get(&result, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindMovies returns the movies
func (r *repo) FindMovies() ([]*Movie, error) {
	var result []*Movie
	err := r.DB.Select(&result, "SELECT * FROM movie ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMovieReviews returns all reviews for a movie
func (r *repo) FindMovieReviews(movieID string) ([]*Review, error) {
	var result []*Review
	err := r.DB.Select(&result, "SELECT * FROM review WHERE movie_id=$1 ORDER BY created_at DESC", movieID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SaveMovie stores a movie
func (r *repo) SaveMovie(movie *Movie) error {
	tx := r.DB.MustBegin()
	tx.NamedExec("INSERT INTO movie (id, title, release_year, duration, short_description, created_at) VALUES (:id, :title, :release_year, :duration, :short_description, :created_at)", &movie)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SaveMovieReview stores a review for a movie
func (r *repo) SaveMovieReview(review *Review) error {
	tx := r.DB.MustBegin()
	tx.NamedExec("INSERT INTO review (id, movie_id, first_name, last_name, score, text, created_at) VALUES (:id, :movie_id, :first_name, :last_name, :score, :text, :created_at)", &review)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
