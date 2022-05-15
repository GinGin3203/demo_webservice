package mongo

import (
	"time"
	"whitesource_home_assingment/internal/movie"
)

type movieDoc struct {
	ID        string    `bson:"_id,omitempty"`
	Title     string    `bson:"title"`
	Year      int       `bson:"year"`
	Country   string    `bson:"country"`
	CreatedAt time.Time `bson:"created_at"`
}

func newMovieDoc(m *movie.Movie) *movieDoc {
	return &movieDoc{
		ID:        m.ID,
		Title:     m.Title,
		Year:      m.Year,
		Country:   m.Country,
		CreatedAt: time.Now(),
	}
}

func (doc *movieDoc) toMovie() *movie.Movie {
	return &movie.Movie{
		Title:     doc.Title,
		Year:      doc.Year,
		Country:   doc.Country,
		ID:        doc.ID,
		CreatedAt: doc.CreatedAt,
	}
}
