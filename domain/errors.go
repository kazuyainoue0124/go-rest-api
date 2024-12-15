package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid = errors.New("invalid input")
)