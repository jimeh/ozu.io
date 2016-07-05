package goleveldbstore

import (
	"errors"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

// DefaultSequenceKey is used by NextSequence().
var DefaultSequenceKey = []byte("__SEQUENCE_ID__")

// ErrNotFound is returned when Get() tries to fetch a non-existent key.
var ErrNotFound = errors.New("not found")

// New creates a new Store using given path to persist data.
func New(path string) (*Store, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return &Store{}, err
	}

	store := Store{
		DB:          db,
		SequenceKey: DefaultSequenceKey,
	}

	return &store, nil
}

// Store allows storing data into a goleveldb database.
type Store struct {
	DB          *leveldb.DB
	SequenceKey []byte
}

// Close underlying goleveldb database.
func (s *Store) Close() error {
	return s.DB.Close()
}

// Get a given key's value.
func (s *Store) Get(key []byte) ([]byte, error) {
	value, err := s.DB.Get(key, nil)
	if err != nil && err.Error() == "leveldb: not found" {
		return nil, ErrNotFound
	}

	return value, err
}

// Set a given key's to the specified value.
func (s *Store) Set(key []byte, value []byte) error {
	return s.DB.Put(key, value, nil)
}

// Delete a given key.
func (s *Store) Delete(key []byte) error {
	return s.DB.Delete(key, nil)
}

// NextSequence returns a auto-incrementing int.
func (s *Store) NextSequence() (int, error) {
	return s.Incr(s.SequenceKey)
}

// Incr increments a given key (must be numeric-like value)
func (s *Store) Incr(key []byte) (int, error) {
	tx, err := s.DB.OpenTransaction()
	if err != nil {
		return -1, err
	}

	value, err := tx.Get(key, nil)
	if value == nil {
		value = []byte("0")
	}

	num, err := strconv.Atoi(string(value))
	if err != nil {
		return -1, err
	}

	num++
	value = []byte(strconv.Itoa(num))

	err = tx.Put(key, value, nil)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return num, nil
}
