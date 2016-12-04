package storage

import "errors"

// ErrNotFound is the default error message when data is not found.
var ErrNotFound = errors.New("not found")

// Store defines a standard interface for storage
type Store interface {
	Close() error
	Create(UID []byte, URL []byte) (*Record, error)
	FindByUID(UID []byte) (*Record, error)
	FindByURL(URL []byte) (*Record, error)
	DeleteByUID(UID []byte) (*Record, error)
	DeleteByURL(URL []byte) (*Record, error)
	NextSequence() (int, error)
}

// Record provides a standard way to refer to a shortened URL.
type Record struct {
	UID []byte
	URL []byte
}
