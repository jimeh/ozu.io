package storage

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
