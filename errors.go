package gokoreanbots

import "errors"

var (
	// ErrBadRequest returned when sent wrong body
	ErrBadRequest = errors.New("400 Bad Request")
	// ErrForbidden returned when
	ErrForbidden = errors.New("403 Forbidden")
	// ErrRateLimited returned when rate limited.
	ErrRateLimited = errors.New("429 Too Many Request")
)
