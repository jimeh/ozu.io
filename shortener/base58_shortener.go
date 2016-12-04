package shortener

import (
	"errors"

	"github.com/jimeh/go-base58"
	"github.com/jimeh/ozu.io/storage"
)

var urlKeyPrefix = []byte("url:")
var uidKeyPrefix = []byte("uid:")
var errInvalidUID = errors.New("invalid UID")

// NewBase58 returns a new *Base58Shortner that uses the given storage.Store.
func NewBase58(store storage.Store) *Base58Shortener {
	return &Base58Shortener{Store: store}
}

// Base58Shortener shortens URLs via base 58 encoding.
type Base58Shortener struct {
	Store storage.Store
}

// Shorten a given URL.
func (s *Base58Shortener) Shorten(rawURL []byte) (*storage.Record, error) {
	url, err := NormalizeURL(rawURL)
	if err != nil {
		return &storage.Record{}, err
	}

	record, err := s.Store.FindByURL(url)
	if err == nil {
		return record, nil
	} else if err != storage.ErrNotFound {
		return &storage.Record{}, err
	}

	uid, err := s.newUID()
	if err != nil {
		return &storage.Record{}, err
	}

	record, err = s.Store.Create(uid, url)
	if err != nil {
		return &storage.Record{}, err
	}

	return record, nil
}

// Lookup the URL of a given UID.
func (s *Base58Shortener) Lookup(uid []byte) (*storage.Record, error) {
	_, err := base58.Decode(uid)
	if err != nil {
		return &storage.Record{}, errInvalidUID
	}

	return s.Store.FindByUID(uid)
}

func (s *Base58Shortener) newUID() ([]byte, error) {
	index, err := s.Store.NextSequence()
	if err != nil {
		return nil, err
	}

	return base58.Encode(index), nil
}
