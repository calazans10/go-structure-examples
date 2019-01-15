package main

// Repository defines the repository interface
type Repository interface {
	FindMovies() ([]*Movie, error)
	FindMovie(id string) (*Movie, error)
	FindReview(movieID string) ([]*Review, error)
	SaveMovie(movie *Movie) error
	SaveReview(review *Review) error
}
