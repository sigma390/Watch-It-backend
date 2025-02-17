package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Genre       string    `json:"genre"`
	Rating      float64   `json:"rating"`
}
