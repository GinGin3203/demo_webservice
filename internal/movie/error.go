package movie

import "errors"

var (
	ErrWrongIDFormat = errors.New("received ID has wrong format")
	ErrNotFound      = errors.New("no such movie")
)
