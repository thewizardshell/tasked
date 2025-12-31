package errors

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
)
