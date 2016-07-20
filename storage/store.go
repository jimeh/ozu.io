package storage

// Store defines a standard interface for storage.
type Store interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) error
	NextSequence() (int, error)
}
