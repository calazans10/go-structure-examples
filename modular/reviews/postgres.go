package reviews

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

// FindMovieReviews returns all reviews for a movie
func (r *repo) FindMovieReviews(movieID string) ([]*Review, error) {
	var result []*Review
	err := r.DB.Select(&result, "SELECT * FROM review WHERE movie_id=$1 ORDER BY created_at DESC", movieID)
	if err != nil {
		return nil, err
	}
	return result, nil
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
