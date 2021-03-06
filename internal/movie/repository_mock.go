// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package movie

import (
	"context"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			AddMovieFunc: func(contextMoqParam context.Context, movie *Movie) (*Movie, error) {
// 				panic("mock out the AddMovie method")
// 			},
// 			DeleteMovieByIDFunc: func(ctx context.Context, id string) error {
// 				panic("mock out the DeleteMovieByID method")
// 			},
// 			GetMovieByIDFunc: func(ctx context.Context, id string) (*Movie, error) {
// 				panic("mock out the GetMovieByID method")
// 			},
// 			UpdateMovieByIDFunc: func(ctx context.Context, id string, model *Movie) error {
// 				panic("mock out the UpdateMovieByID method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// AddMovieFunc mocks the AddMovie method.
	AddMovieFunc func(contextMoqParam context.Context, movie *Movie) (*Movie, error)

	// DeleteMovieByIDFunc mocks the DeleteMovieByID method.
	DeleteMovieByIDFunc func(ctx context.Context, id string) error

	// GetMovieByIDFunc mocks the GetMovieByID method.
	GetMovieByIDFunc func(ctx context.Context, id string) (*Movie, error)

	// UpdateMovieByIDFunc mocks the UpdateMovieByID method.
	UpdateMovieByIDFunc func(ctx context.Context, id string, model *Movie) error

	// calls tracks calls to the methods.
	calls struct {
		// AddMovie holds details about calls to the AddMovie method.
		AddMovie []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Movie is the movie argument value.
			Movie *Movie
		}
		// DeleteMovieByID holds details about calls to the DeleteMovieByID method.
		DeleteMovieByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetMovieByID holds details about calls to the GetMovieByID method.
		GetMovieByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// UpdateMovieByID holds details about calls to the UpdateMovieByID method.
		UpdateMovieByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Model is the model argument value.
			Model *Movie
		}
	}
	lockAddMovie        sync.RWMutex
	lockDeleteMovieByID sync.RWMutex
	lockGetMovieByID    sync.RWMutex
	lockUpdateMovieByID sync.RWMutex
}

// AddMovie calls AddMovieFunc.
func (mock *RepositoryMock) AddMovie(contextMoqParam context.Context, movie *Movie) (*Movie, error) {
	if mock.AddMovieFunc == nil {
		panic("RepositoryMock.AddMovieFunc: method is nil but Repository.AddMovie was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Movie           *Movie
	}{
		ContextMoqParam: contextMoqParam,
		Movie:           movie,
	}
	mock.lockAddMovie.Lock()
	mock.calls.AddMovie = append(mock.calls.AddMovie, callInfo)
	mock.lockAddMovie.Unlock()
	return mock.AddMovieFunc(contextMoqParam, movie)
}

// AddMovieCalls gets all the calls that were made to AddMovie.
// Check the length with:
//     len(mockedRepository.AddMovieCalls())
func (mock *RepositoryMock) AddMovieCalls() []struct {
	ContextMoqParam context.Context
	Movie           *Movie
} {
	var calls []struct {
		ContextMoqParam context.Context
		Movie           *Movie
	}
	mock.lockAddMovie.RLock()
	calls = mock.calls.AddMovie
	mock.lockAddMovie.RUnlock()
	return calls
}

// DeleteMovieByID calls DeleteMovieByIDFunc.
func (mock *RepositoryMock) DeleteMovieByID(ctx context.Context, id string) error {
	if mock.DeleteMovieByIDFunc == nil {
		panic("RepositoryMock.DeleteMovieByIDFunc: method is nil but Repository.DeleteMovieByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteMovieByID.Lock()
	mock.calls.DeleteMovieByID = append(mock.calls.DeleteMovieByID, callInfo)
	mock.lockDeleteMovieByID.Unlock()
	return mock.DeleteMovieByIDFunc(ctx, id)
}

// DeleteMovieByIDCalls gets all the calls that were made to DeleteMovieByID.
// Check the length with:
//     len(mockedRepository.DeleteMovieByIDCalls())
func (mock *RepositoryMock) DeleteMovieByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteMovieByID.RLock()
	calls = mock.calls.DeleteMovieByID
	mock.lockDeleteMovieByID.RUnlock()
	return calls
}

// GetMovieByID calls GetMovieByIDFunc.
func (mock *RepositoryMock) GetMovieByID(ctx context.Context, id string) (*Movie, error) {
	if mock.GetMovieByIDFunc == nil {
		panic("RepositoryMock.GetMovieByIDFunc: method is nil but Repository.GetMovieByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetMovieByID.Lock()
	mock.calls.GetMovieByID = append(mock.calls.GetMovieByID, callInfo)
	mock.lockGetMovieByID.Unlock()
	return mock.GetMovieByIDFunc(ctx, id)
}

// GetMovieByIDCalls gets all the calls that were made to GetMovieByID.
// Check the length with:
//     len(mockedRepository.GetMovieByIDCalls())
func (mock *RepositoryMock) GetMovieByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetMovieByID.RLock()
	calls = mock.calls.GetMovieByID
	mock.lockGetMovieByID.RUnlock()
	return calls
}

// UpdateMovieByID calls UpdateMovieByIDFunc.
func (mock *RepositoryMock) UpdateMovieByID(ctx context.Context, id string, model *Movie) error {
	if mock.UpdateMovieByIDFunc == nil {
		panic("RepositoryMock.UpdateMovieByIDFunc: method is nil but Repository.UpdateMovieByID was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		ID    string
		Model *Movie
	}{
		Ctx:   ctx,
		ID:    id,
		Model: model,
	}
	mock.lockUpdateMovieByID.Lock()
	mock.calls.UpdateMovieByID = append(mock.calls.UpdateMovieByID, callInfo)
	mock.lockUpdateMovieByID.Unlock()
	return mock.UpdateMovieByIDFunc(ctx, id, model)
}

// UpdateMovieByIDCalls gets all the calls that were made to UpdateMovieByID.
// Check the length with:
//     len(mockedRepository.UpdateMovieByIDCalls())
func (mock *RepositoryMock) UpdateMovieByIDCalls() []struct {
	Ctx   context.Context
	ID    string
	Model *Movie
} {
	var calls []struct {
		Ctx   context.Context
		ID    string
		Model *Movie
	}
	mock.lockUpdateMovieByID.RLock()
	calls = mock.calls.UpdateMovieByID
	mock.lockUpdateMovieByID.RUnlock()
	return calls
}
