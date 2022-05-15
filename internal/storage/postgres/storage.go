package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"whitesource_home_assingment/internal/movie"
	"whitesource_home_assingment/pkg/must"
)

// a convenience to make sure that Storage implements the desired interface
var _ movie.Repository = &Storage{}

type Storage struct {
	Conn *pgx.Conn
}

func NewStorage(ctx context.Context, url string) *Storage {
	conn := must.NotFail(pgx.Connect(ctx, url))
	must.NotFail(conn.Exec(
		ctx,
		`CREATE TABLE IF NOT EXISTS movies (
				id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
				title varchar(255) NOT NULL,
				country varchar(255) NOT NULL,
				year integer NOT NULL,
				created_at timestamp DEFAULT now()
				)`,
	))
	return &Storage{Conn: conn}
}
func (s *Storage) AddMovie(ctx context.Context, m *movie.Movie) (*movie.Movie, error) {
	row := s.Conn.QueryRow(
		ctx,
		`INSERT INTO movies (title, country, year) VALUES ($1, $2, $3) RETURNING 
				id, created_at`,
		m.Title, m.Country, m.Year,
	)
	err := row.Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Storage) UpdateMovieByID(ctx context.Context, id string, m *movie.Movie) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return movie.ErrWrongIDFormat
	}
	_, err = s.Conn.Exec(
		ctx, `UPDATE movies SET title=$1, country=$2, year=$3 WHERE id=$4`,
		m.Title, m.Country, m.Year, u,
	)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return movie.ErrNotFound
	default:
		return err
	}
}

func (s *Storage) GetMovieByID(ctx context.Context, id string) (*movie.Movie, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, movie.ErrWrongIDFormat
	}
	row := s.Conn.QueryRow(
		ctx,
		`SELECT id, title, country, year, created_at FROM movies WHERE id=$1`, u,
	)
	var m movie.Movie
	err = row.Scan(&m.ID, &m.Title, &m.Country, &m.Year, &m.CreatedAt)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, movie.ErrNotFound
	case errors.Is(err, nil):
		return &m, nil
	default:
		return nil, err
	}
}

func (s *Storage) DeleteMovieByID(ctx context.Context, id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return movie.ErrWrongIDFormat
	}
	_, err = s.Conn.Exec(
		ctx,
		`DELETE FROM movies WHERE id=$1`, u,
	)
	if err != nil {
		return err
	}
	return nil
}
