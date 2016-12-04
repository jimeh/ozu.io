package goleveldbstore

import (
	"strconv"

	"github.com/jimeh/ozu.io/storage"
	"github.com/syndtr/goleveldb/leveldb"
)

const errLeveldbNotFound = "leveldb: not found"

// DefaultSequenceKey is used by NextSequence().
var DefaultSequenceKey = []byte("__SEQUENCE_ID__")

var uidKeyPrefix = []byte("!")
var urlKeyPrefix = []byte("#")

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

// Create a given Record.
func (s *Store) Create(uid []byte, url []byte) (*storage.Record, error) {
	tx, err := s.DB.OpenTransaction()
	if err != nil {
		return &storage.Record{}, err
	}

	err = tx.Put(s.uidKey(uid), url, nil)
	if err != nil {
		return &storage.Record{}, err
	}

	err = tx.Put(s.urlKey(url), uid, nil)
	if err != nil {
		return &storage.Record{}, err
	}

	err = tx.Commit()
	if err != nil {
		return &storage.Record{}, err
	}

	return &storage.Record{UID: uid, URL: url}, nil
}

// FindByUID looks up records based on their UID.
func (s *Store) FindByUID(uid []byte) (*storage.Record, error) {
	value, err := s.DB.Get(s.uidKey(uid), nil)
	if err != nil {
		if err.Error() == errLeveldbNotFound {
			return &storage.Record{}, storage.ErrNotFound
		}
		return &storage.Record{}, err
	}

	return &storage.Record{UID: uid, URL: value}, nil
}

// FindByURL looks up records based on their URL.
func (s *Store) FindByURL(url []byte) (*storage.Record, error) {
	value, err := s.DB.Get(s.urlKey(url), nil)
	if err != nil {
		if err.Error() == errLeveldbNotFound {
			return &storage.Record{}, storage.ErrNotFound
		}
		return &storage.Record{}, err
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

	err = s.delete(record)
	if err != nil {
		return &storage.Record{}, err
	}

	return record, nil
}

func (s *Store) delete(r *storage.Record) error {
	tx, err := s.DB.OpenTransaction()
	if err != nil {
		return err
	}

	err = tx.Delete(s.uidKey(r.UID), nil)
	if err != nil && err.Error() == errLeveldbNotFound {
		return err
	}

	err = tx.Delete(s.urlKey(r.URL), nil)
	if err != nil && err.Error() == errLeveldbNotFound {
		return err
	}

	return tx.Commit()
}

func (s *Store) uidKey(uid []byte) []byte {
	return append(uidKeyPrefix, uid...)
}

func (s *Store) urlKey(url []byte) []byte {
	return append(urlKeyPrefix, url...)
}

// NextSequence returns a auto-incrementing int.
func (s *Store) NextSequence() (int, error) {
	return s.incr(s.SequenceKey)
}

// Incr increments a given key (must be numeric-like value)
func (s *Store) incr(key []byte) (int, error) {
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
