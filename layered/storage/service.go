package storage

import (
	"time"

	"github.com/calazans10/go-structure-examples/layered/models"
	"github.com/gofrs/uuid"
)

// Service defines the service structure
type Service struct {
	repo Repository
}

// NewService creates a new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// FindMovie returns a movie
func (s *Service) FindMovie(id string) (*models.Movie, error) {
	return s.repo.FindMovie(id)
}

// FindMovies returns the movies
func (s *Service) FindMovies() ([]*models.Movie, error) {
	return s.repo.FindMovies()
}

// FindMovieReviews returns all reviews for a movie
func (s *Service) FindMovieReviews(movieID string) ([]*models.Review, error) {
	return s.repo.FindMovieReviews(movieID)
}

// SaveMovie stores a movie
func (s *Service) SaveMovie(movie *models.Movie) error {
	uuid, _ := uuid.NewV4()
	movie.ID = uuid.String()
	movie.CreatedAt = time.Now()
	return s.repo.SaveMovie(movie)
}

// SaveMovieReview stores a review for a movie
func (s *Service) SaveMovieReview(review *models.Review) error {
	uuid, _ := uuid.NewV4()
	review.ID = uuid.String()
	review.CreatedAt = time.Now()
	return s.repo.SaveMovieReview(review)
}
