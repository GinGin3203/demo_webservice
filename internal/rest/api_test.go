package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"whitesource_home_assingment/internal/movie"
	"whitesource_home_assingment/pkg/must"
)

var (
	idNotFound    = "idNotFound"
	idWrongFormat = "idWrongFormat"
	idDummyMovie  = "idDummyMovie"
	idExists      = "idExists"

	dummyMovie = &movie.Movie{
		Title:     "dummy",
		Year:      0,
		Country:   "dummy",
		ID:        idDummyMovie,
		CreatedAt: time.Unix(0, 0),
	}

	dummyMovieMapped = gin.H{
		"title":      "dummy",
		"year":       float64(0),
		"country":    "dummy",
		"id":         idDummyMovie,
		"created_at": "1970-01-01T03:00:00+03:00",
	}
)

func setup(repo *movie.RepositoryMock) *gin.Engine {
	service := movie.Service{
		MovieRepository: repo,
	}
	return (&Server{MovieService: service}).NewRoutesV1(gin.Default())
}

func TestGet(t *testing.T) {
	router := setup(
		&movie.RepositoryMock{
			GetMovieByIDFunc: func(ctx context.Context, id string) (*movie.Movie, error) {
				switch id {
				case idNotFound:
					return nil, movie.ErrNotFound
				case idWrongFormat:
					return nil, movie.ErrWrongIDFormat
				case idDummyMovie:
					return dummyMovie, nil
				default:
					panic("unexpected value")
				}
			},
		},
	)

	for name, tc := range map[string]struct {
		path         string
		expectedCode int
		bodyExpected any
	}{
		"IDNotFound": {
			path:         "/v1/movie/" + idNotFound,
			expectedCode: http.StatusNotFound,
			bodyExpected: gin.H{"message": "no such movie"},
		},
		"IDWrongFormat": {
			path:         "/v1/movie/" + idWrongFormat,
			expectedCode: http.StatusNotFound,
			bodyExpected: gin.H{"message": "no such movie"},
		},
		"IDDummyMovie": {
			path:         "/v1/movie/" + idDummyMovie,
			expectedCode: http.StatusOK,
			bodyExpected: dummyMovieMapped,
		},
	} {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := must.NotFail(http.NewRequest(http.MethodGet, tc.path, nil))
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			var dest gin.H
			must.NoError(json.Unmarshal(must.NotFail(ioutil.ReadAll(w.Body)), &dest))
			assert.EqualValues(t, tc.bodyExpected, dest)
		})
	}
}

func TestPost(t *testing.T) {
	// we only check that POST /v1/movie returns JSON containing the same
	// `title`, `year`, `country` that we sent in request body
	router := setup(
		&movie.RepositoryMock{
			AddMovieFunc: func(ctx context.Context, m *movie.Movie) (*movie.Movie, error) {
				return m, nil
			},
		},
	)
	w := httptest.NewRecorder()
	body := bytes.NewBuffer(must.NotFail(json.Marshal(dummyMovieMapped)))
	req := must.NotFail(http.NewRequest(http.MethodPost, "/v1/movie", body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var dest gin.H
	must.NoError(json.Unmarshal(must.NotFail(ioutil.ReadAll(w.Body)), &dest))

	assert.EqualValues(t, dummyMovieMapped["title"], dest["title"])
	assert.EqualValues(t, dummyMovieMapped["year"], dest["year"])
	assert.EqualValues(t, dummyMovieMapped["country"], dest["country"])
}

func TestPut(t *testing.T) {
	router := setup(
		&movie.RepositoryMock{
			UpdateMovieByIDFunc: func(ctx context.Context, id string, _ *movie.Movie) error {
				switch id {
				case idNotFound:
					return movie.ErrNotFound
				case idWrongFormat:
					return movie.ErrWrongIDFormat
				case idExists:
					return nil
				default:
					panic("unexpected value")
				}
			},
		},
	)

	for name, tc := range map[string]struct {
		path         string
		expectedCode int
		bodyExpected any
	}{
		"IDNotFound": {
			path:         "/v1/movie/" + idNotFound,
			expectedCode: http.StatusNotFound,
			bodyExpected: gin.H{"message": "no such movie"},
		},
		"IDWrongFormat": {
			path:         "/v1/movie/" + idWrongFormat,
			expectedCode: http.StatusNotFound,
			bodyExpected: gin.H{"message": "no such movie"},
		},
		"IDExists": {
			path:         "/v1/movie/" + idExists,
			expectedCode: http.StatusOK,
			bodyExpected: msgOK,
		},
	} {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			body := bytes.NewBuffer(must.NotFail(json.Marshal(dummyMovieMapped)))
			req := must.NotFail(http.NewRequest(http.MethodPut, tc.path, body))
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			var dest gin.H
			must.NoError(json.Unmarshal(must.NotFail(ioutil.ReadAll(w.Body)), &dest))
			assert.EqualValues(t, tc.bodyExpected, dest)
		})
	}
}

func TestDelete(t *testing.T) {
	router := setup(
		&movie.RepositoryMock{
			DeleteMovieByIDFunc: func(ctx context.Context, id string) error {
				switch id {
				case idExists:
					return nil
				case idWrongFormat:
					return movie.ErrWrongIDFormat
				case idNotFound:
					return movie.ErrNotFound
				default:
					panic("unexpected value")
				}
			},
		},
	)

	for name, tc := range map[string]struct {
		path         string
		expectedCode int
	}{
		"IDNotFound": {
			path:         "/v1/movie/" + idNotFound,
			expectedCode: http.StatusNoContent,
		},
		"IDWrongFormat": {
			path:         "/v1/movie/" + idWrongFormat,
			expectedCode: http.StatusNoContent,
		},
		"IDExists": {
			path:         "/v1/movie/" + idExists,
			expectedCode: http.StatusNoContent,
		},
	} {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := must.NotFail(http.NewRequest(http.MethodDelete, tc.path, nil))
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
