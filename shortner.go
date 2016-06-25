package main

import (
	"log"
	"strconv"

	"github.com/jbenet/go-base58"
	"github.com/syndtr/goleveldb/leveldb"
)

// NewShortner returns a new Shortner.
func NewShortner() *Shortner {
	db, err := leveldb.OpenFile("shortner.db", nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Shortner{
		Db:        db,
		KeyPrefix: []byte("url:"),
	}
}

// Shortner shortens URLs.
type Shortner struct {
	Db        *leveldb.DB
	KeyPrefix []byte
}

// Close calls Close() on the goleveldb.
func (s *Shortner) Close() {
	err := s.Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Get reads from the database.
func (s *Shortner) Get(key []byte) []byte {
	value, _ := s.Db.Get(key, nil)
	return value
}

// Set writes to the database.
func (s *Shortner) Set(key []byte, value []byte) error {
	return s.Db.Put(key, value, nil)
}

// Shorten shortens given URL
func (s *Shortner) Shorten(url []byte) ([]byte, error) {
	uid, err := s.newUID()
	if err != nil {
		return []byte{}, err
	}

	key := s.makeKey(uid)
	err = s.Set(key, url)
	if err != nil {
		return []byte{}, err
	}

	return uid, nil
}

// Lookup attempts to fetch the value for given UID
func (s *Shortner) Lookup(uid []byte) []byte {
	return s.Get(s.makeKey(uid))
}

func (s *Shortner) makeKey(uid []byte) []byte {
	return append(s.KeyPrefix, uid...)
}

func (s *Shortner) newUID() ([]byte, error) {
	key := []byte("index")

	tx, err := s.Db.OpenTransaction()
	defer tx.Commit()

	if err != nil {
		return []byte{}, err
	}

	index, _ := tx.Get(key, nil)
	if index == nil {
		index = []byte("1")
	}

	num, _ := strconv.Atoi(string(index))
	num++
	index = []byte(strconv.Itoa(num))

	err = tx.Put(key, index, nil)
	if err != nil {
		return []byte{}, err
	}

	uid := base58.Encode(index)
	return []byte(uid), nil
}
