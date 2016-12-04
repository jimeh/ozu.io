package shortener

import "github.com/jimeh/ozu.io/storage"

// Shortener defines a shortener interface for shortening URLs.
type Shortener interface {
	Shorten([]byte) (*storage.Record, error)
	Lookup([]byte) (*storage.Record, error)
}
