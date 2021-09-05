package gokoreanbots

import (
	"errors"
	"fmt"
)

func getStatusError(statusCode int) error {
	switch statusCode {
	case 400:
		return ErrBadRequest
	case 403:
		return ErrForbidden
	case 429:
		return ErrRateLimited
	case 200:
		return nil
	default:
		return errors.New(fmt.Sprintf("uncaught status code %d", statusCode))
	}
}
