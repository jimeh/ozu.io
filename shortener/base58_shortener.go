package shortener

import (
	"crypto/sha1"
	"fmt"

	"github.com/jimeh/go-base58"
	"github.com/jimeh/ozu.io/storage"
)

var urlKeyPrefix = []byte("url:")
var uidKeyPrefix = []byte("uid:")

// NewBase58 returns a new *Base58Shortner that uses the given storage.Store.
func NewBase58(store storage.Store) *Base58Shortener {
	return &Base58Shortener{Store: store}
}

// Base58Shortener shortens URLs via base 58 encoding.
type Base58Shortener struct {
	Store storage.Store
}

// Shorten a given URL.
func (s *Base58Shortener) Shorten(rawURL []byte) (uid []byte, url []byte, err error) {
	url, err = NormalizeURL(rawURL)
	if err != nil {
		return nil, nil, err
	}

	urlKey := s.makeURLKey(url)
	uid, err = s.Store.Get(urlKey)

	if uid != nil && err == nil {
		return uid, url, nil
	} else if err != nil && err.Error() != "not found" {
		return nil, nil, err
	}

	uid, err = s.newUID()
	if err != nil {
		return nil, nil, err
	}

	err = s.Store.Set(urlKey, uid)
	if err != nil {
		return nil, nil, err
	}

	uidKey := s.makeUIDKey(uid)
	err = s.Store.Set(uidKey, url)
	if err != nil {
		return nil, nil, err
	}

	return uid, url, nil
}

// Lookup the URL of a given UID.
func (s *Base58Shortener) Lookup(uid []byte) ([]byte, error) {
	uidKey := s.makeUIDKey(uid)

	url, err := s.Store.Get(uidKey)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *Base58Shortener) newUID() ([]byte, error) {
	index, err := s.Store.NextSequence()
	if err != nil {
		return nil, err
	}

	return base58.Encode(index), nil
}

func (s *Base58Shortener) makeUIDKey(uid []byte) []byte {
	return append(uidKeyPrefix, uid...)
}

func (s *Base58Shortener) makeURLKey(rawURL []byte) []byte {
	urlSHA := fmt.Sprintf("%x", sha1.Sum(rawURL))
	return append(urlKeyPrefix, urlSHA...)
}
