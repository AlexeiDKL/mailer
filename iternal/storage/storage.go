package storage

import "errors"

type Pair struct {
	Type  string
	Value any
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
