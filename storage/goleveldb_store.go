package storage

import (
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

// GoleveldbStore allows storing data into a goleveldb database.
type GoleveldbStore struct {
	DB *leveldb.DB
}

// NewGoleveldbStore creates a new GoleveldbStore using given path to persist
// data.
func NewGoleveldbStore(path string) (GoleveldbStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return GoleveldbStore{}, err
	}

	return GoleveldbStore{DB: db}, nil
}

// Close underlying goleveldb database.
func (s *GoleveldbStore) Close() error {
	return s.DB.Close()
}

// Get a given key's value.
func (s *GoleveldbStore) Get(key []byte) ([]byte, error) {
	value, err := s.DB.Get(key, nil)
	if err != nil && err.Error() == "leveldb: not found" {
		return value, nil
	}

	return value, err
}

// Set a given key's to the specified value.
func (s *GoleveldbStore) Set(key []byte, value []byte) error {
	return s.DB.Put(key, value, nil)
}

// Delete a given key.
func (s *GoleveldbStore) Delete(key []byte) error {
	return s.DB.Delete(key, nil)
}

// Incr increments a given key (must be numeric-like value)
func (s *GoleveldbStore) Incr(key []byte) (int, error) {
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
