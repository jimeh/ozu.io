package shortener

import "github.com/jimeh/ozu.io/storage"

// New returns a new *Base58Shortner that uses the given storage.Store.
func New(store storage.Store) Shortener {
	return NewBase58(store)
}
