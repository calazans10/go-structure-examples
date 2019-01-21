package movies

// Repository defines the repository interface
type Repository interface {
	FindMovie(id string) (*Movie, error)
	FindMovies() ([]*Movie, error)
	SaveMovie(movie *Movie) error
}
