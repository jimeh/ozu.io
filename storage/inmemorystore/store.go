package inmemorystore

import (
	"sync"

	"github.com/jimeh/ozu.io/storage"
)

// New creates a new Store using given path to persist data.
func New() (*Store, error) {
	store := &Store{
		UIDMap:   map[string][]byte{},
		URLMap:   map[string][]byte{},
		Sequence: 0,
	}
	return store, nil
}

// Store allows storing data into a in-memory map.
type Store struct {
	sync.RWMutex
	UIDMap   map[string][]byte
	URLMap   map[string][]byte
	Sequence int
}

// Close database.
func (s *Store) Close() error {
	s.Lock()
	s.UIDMap = make(map[string][]byte)
	s.URLMap = make(map[string][]byte)
	s.Sequence = 0
	s.Unlock()
	return nil
}

// Create a given Record
func (s *Store) Create(uid []byte, url []byte) (*storage.Record, error) {
	s.Lock()
	s.UIDMap[string(uid)] = url
	s.URLMap[string(url)] = uid
	s.Unlock()
	return &storage.Record{UID: uid, URL: url}, nil
}

// FindByUID looks up records based on their UID.
func (s *Store) FindByUID(uid []byte) (*storage.Record, error) {
	s.RLock()
	value := s.UIDMap[string(uid)]
	s.RUnlock()
	if value == nil {
		return &storage.Record{}, storage.ErrNotFound
	}
	return &storage.Record{UID: uid, URL: value}, nil
}

// FindByURL looks up records based on their URL.
func (s *Store) FindByURL(url []byte) (*storage.Record, error) {
	s.RLock()
	value := s.URLMap[string(url)]
	s.RUnlock()
	if value == nil {
		return &storage.Record{}, storage.ErrNotFound
	}
	return &storage.Record{UID: value, URL: url}, nil
}

// DeleteByUID deletes records based on their UID.
func (s *Store) DeleteByUID(uid []byte) (*storage.Record, error) {
	record, err := s.FindByUID(uid)
	if err != nil {
		return &storage.Record{}, err
	}

	s.delete(record)
	return record, nil
}

// DeleteByURL deletes records based on their URL.
func (s *Store) DeleteByURL(url []byte) (*storage.Record, error) {
	record, err := s.FindByURL(url)
	if err != nil {
		return &storage.Record{}, err
	}

	s.delete(record)
	return record, nil
}

func (s *Store) delete(r *storage.Record) {
	s.Lock()
	delete(s.UIDMap, string(r.UID))
	delete(s.URLMap, string(r.URL))
	s.Unlock()
}

// NextSequence returns a auto-incrementing int.
func (s *Store) NextSequence() (int, error) {
	s.Lock()
	s.Sequence++
	seq := s.Sequence
	s.Unlock()
	return seq, nil
}
