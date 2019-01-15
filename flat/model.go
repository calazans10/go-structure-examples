package main

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

// Review defines the properties of a movie review
type Review struct {
	ID        string    `db:"id" json:"id"`
	MovieID   string    `db:"movie_id" json:"movie_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Score     int       `db:"score" json:"score"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
