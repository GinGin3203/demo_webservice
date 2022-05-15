package rest

import (
	"time"
	"whitesource_home_assingment/internal/movie"
)

type movieJSON struct {
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Country   string    `json:"country"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func newMovieJSON(m *movie.Movie) *movieJSON {
	return &movieJSON{
		Title:     m.Title,
		Year:      m.Year,
		Country:   m.Country,
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
	}

}

func (dto *movieJSON) toMovie() *movie.Movie {
	return &movie.Movie{
		Title:   dto.Title,
		Year:    dto.Year,
		Country: dto.Country,
	}
}
