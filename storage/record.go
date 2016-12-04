package storage

// Record provides a standard way to refer to a shortened URL.
type Record struct {
	UID []byte
	URL []byte
}
