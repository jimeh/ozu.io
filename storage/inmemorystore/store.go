package inmemorystore

import (
	"errors"
	"sync"
)

// ErrNotFound is returned when Get() tries to fetch a non-existent key.
var ErrNotFound = errors.New("not found")

// New creates a new Store using given path to persist data.
func New() (*Store, error) {
	store := &Store{
		Data:     map[string][]byte{},
		Sequence: 0,
	}
	return store, nil
}

// Store allows storing data into a in-memory map.
type Store struct {
	sync.RWMutex
	Data     map[string][]byte
	Sequence int
}

// Get a given key's value.
func (s *Store) Get(key []byte) ([]byte, error) {
	s.RLock()
	value := s.Data[string(key)]
	s.RUnlock()
	if value == nil {
		return nil, ErrNotFound
	}
	return value, nil
}

// Set a given key's to the specified value.
func (s *Store) Set(key []byte, value []byte) error {
	s.Lock()
	s.Data[string(key)] = value
	s.Unlock()
	return nil
}

// Delete a given key.
func (s *Store) Delete(key []byte) error {
	s.Lock()
	delete(s.Data, string(key))
	s.Unlock()
	return nil
}

// NextSequence returns a auto-incrementing int.
func (s *Store) NextSequence() (int, error) {
	s.Lock()
	s.Sequence++
	seq := s.Sequence
	s.Unlock()
	return seq, nil
}
