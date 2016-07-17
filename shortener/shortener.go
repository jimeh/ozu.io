package shortener

// Shortener defines a shortener interface for shortening URLs.
type Shortener interface {
	Shorten([]byte) ([]byte, []byte, error)
	Lookup([]byte) ([]byte, error)
}
