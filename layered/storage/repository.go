package storage

import "github.com/calazans10/go-structure-examples/layered/models"

// Repository defines the repository interface
type Repository interface {
	FindMovie(id string) (*models.Movie, error)
	FindMovies() ([]*models.Movie, error)
	FindMovieReviews(movieID string) ([]*models.Review, error)
	SaveMovie(movie *models.Movie) error
	SaveMovieReview(review *models.Review) error
}
