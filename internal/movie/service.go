package movie

import (
	"context"
	"time"
)

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	AddMovie(context.Context, *Movie) (*Movie, error)
	UpdateMovieByID(ctx context.Context, id string, model *Movie) error
	GetMovieByID(ctx context.Context, id string) (*Movie, error)
	DeleteMovieByID(ctx context.Context, id string) error
}

type Service struct {
	MovieRepository Repository
}

type Movie struct {
	Title     string
	Year      int
	Country   string
	ID        string
	CreatedAt time.Time
}

func (s *Service) Add(ctx context.Context, m *Movie) (*Movie, error) {
	return s.MovieRepository.AddMovie(ctx, m)
}
func (s *Service) Update(ctx context.Context, id string, m *Movie) error {
	return s.MovieRepository.UpdateMovieByID(ctx, id, m)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.MovieRepository.DeleteMovieByID(ctx, id)
}
func (s *Service) GetOne(ctx context.Context, id string) (*Movie, error) {
	return s.MovieRepository.GetMovieByID(ctx, id)
}
