package main

type storage struct {
}

func (s *storage) FindMovies() ([]*Movie, error) {
}

func (s *storage) FindMovie(id string) (*Movie, error) {
}

func (s *storage) FindReview(movieID string) ([]*Review, error) {
}

func (s *storage) SaveMovie(movie *Movie) error {
}

func (s *storage) SaveReview(review *Review) error {
}
