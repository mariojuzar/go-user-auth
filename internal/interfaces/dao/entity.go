package dao

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrNilParam = errors.New("invalid nil param")
)
