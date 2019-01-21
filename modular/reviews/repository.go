package reviews

// Repository defines the repository interface
type Repository interface {
	FindMovieReviews(movieID string) ([]*Review, error)
	SaveMovieReview(review *Review) error
}
