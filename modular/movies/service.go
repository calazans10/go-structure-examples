package movies

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

// FindMovie returns a movie
func (s *Service) FindMovie(id string) (*Movie, error) {
	return s.repo.FindMovie(id)
}

// FindMovies returns the movies
func (s *Service) FindMovies() ([]*Movie, error) {
	return s.repo.FindMovies()
}

// SaveMovie stores a movie
func (s *Service) SaveMovie(movie *Movie) error {
	uuid, _ := uuid.NewV4()
	movie.ID = uuid.String()
	movie.CreatedAt = time.Now()
	return s.repo.SaveMovie(movie)
}
