package gokoreanbots

import "errors"

var (
	// ErrBadRequest returned when sent wrong body
	ErrBadRequest = errors.New("400 Bad Request")
	// ErrForbidden returned when get wrong Authorization header
	ErrForbidden = errors.New("403 Forbidden")
	// ErrRateLimited returned when rate limited.
	ErrRateLimited = errors.New("429 Too Many Request")
	// ErrUnknownStatusCode returned when status code is not 200 or not defined as error
	ErrUnknownStatusCode = errors.New("api returned unknown status code")
)
