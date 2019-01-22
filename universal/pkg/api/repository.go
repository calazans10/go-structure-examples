package api

// Repository defines the repository interface
type Repository interface {
	FindMovie(id string) (*Movie, error)
	FindMovies() ([]*Movie, error)
	FindMovieReviews(movieID string) ([]*Review, error)
	SaveMovie(movie *Movie) error
	SaveMovieReview(review *Review) error
}
