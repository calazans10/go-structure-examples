package reviews

import (
	"time"

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

// FindMovieReviews returns all reviews for a movie
func (s *Service) FindMovieReviews(movieID string) ([]*Review, error) {
	return s.repo.FindMovieReviews(movieID)
}

// SaveMovieReview stores a review for a movie
func (s *Service) SaveMovieReview(review *Review) error {
	uuid, _ := uuid.NewV4()
	review.ID = uuid.String()
	review.CreatedAt = time.Now()
	return s.repo.SaveMovieReview(review)
}
