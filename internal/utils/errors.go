package utils

import "errors"

var (
	ErrInvalidID  = errors.New("invalid ID format")
	ErrNotFound   = errors.New("record not found")
	ErrBadRequest = errors.New("bad request data")
	ErrInternal   = errors.New("internal server error")
)
