package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"whitesource_home_assingment/internal/movie"
	"whitesource_home_assingment/pkg/must"
)

type Storage struct {
	movies *mongo.Collection
}

// a convenience to make sure that Storage implements the desired interface
var _ movie.Repository = &Storage{}

func NewStorage(ctx context.Context, uri string) *Storage {
	c := must.NotFail(mongo.Connect(ctx, options.Client().ApplyURI(uri)))
	must.NoError(c.Ping(ctx, nil))
	col := c.
		Database("myapp").
		Collection("movies")
	return &Storage{
		movies: col,
	}
}

func (s *Storage) AddMovie(ctx context.Context, m *movie.Movie) (*movie.Movie, error) {
	doc := newMovieDoc(m)
	res, err := s.movies.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	doc.ID = id
	return doc.toMovie(), nil
}

type updatedFields struct {
	Title   string `bson:"title"`
	Year    int    `bson:"year"`
	Country string `bson:"country"`
}

func (s *Storage) UpdateMovieByID(ctx context.Context, hex string, m *movie.Movie) error {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return fmt.Errorf("%w: %s", movie.ErrWrongIDFormat, hex)
	}
	_, err = s.movies.UpdateByID(
		ctx,
		objID,
		bson.D{{"$set",
			&updatedFields{
				Title:   m.Title,
				Year:    m.Year,
				Country: m.Country,
			},
		},
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMovieByID(ctx context.Context, hex string) (*movie.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", movie.ErrWrongIDFormat, hex)
	}
	res := s.movies.FindOne(ctx, bson.D{{"_id", objID}})
	if res.Err() == mongo.ErrNoDocuments {
		return nil, movie.ErrNotFound
	} else if res.Err() != nil {
		return nil, res.Err()
	}
	var m movieDoc
	err = res.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m.toMovie(), nil
}

func (s *Storage) DeleteMovieByID(ctx context.Context, hex string) error {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return fmt.Errorf("%w: %s", movie.ErrWrongIDFormat, hex)
	}
	_, err = s.movies.DeleteOne(ctx, bson.D{{"_id", objID}})
	if err != nil {
		return err
	}
	return nil
}
