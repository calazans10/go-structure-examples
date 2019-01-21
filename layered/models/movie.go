package models

import "time"

// Movie defines the properties of a movie
type Movie struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	ReleaseYear int       `db:"release_year" json:"release_year"`
	Duration    int       `db:"duration" json:"duration"`
	ShortDesc   string    `db:"short_description" json:"short_description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
