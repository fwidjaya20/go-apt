package httpError

import "net/http"

// ConflictError is a struct to contain bad request http error
type ConflictError struct {
	BaseError
}

// Conflict is a constructor to create ConflictError instance
func Conflict(err interface{}) error {
	return &ConflictError{
		BaseError: Code(http.StatusConflict, err),
	}
}

// IsConflictError check whether given error is a ConflictError
func IsConflictError(err error) bool {
	cause := getCause(err)
	_, ok := cause.(*ConflictError)
	return ok
}
