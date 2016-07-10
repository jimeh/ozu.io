package shortner

import (
	"github.com/jimeh/go-base58"
	"github.com/jimeh/ozu.io/storage"
)

// New returns a new *Shortner that uses the given storage.Store.
func New(store storage.Store) *Shortner {
	return &Shortner{Store: store}
}

var urlKeyPrefix = []byte("url:")
var uidKeyPrefix = []byte("uid:")

// Shortner interface
type Shortner struct {
	Store storage.Store
}

// Shorten a given URL.
func (s *Shortner) Shorten(rawURL []byte) (uid []byte, url []byte, err error) {
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
func (s *Shortner) Lookup(uid []byte) ([]byte, error) {
	uidKey := s.makeUIDKey(uid)

	url, err := s.Store.Get(uidKey)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *Shortner) newUID() ([]byte, error) {
	index, err := s.Store.NextSequence()
	if err != nil {
		return nil, err
	}

	return base58.Encode(index), nil
}

func (s *Shortner) makeUIDKey(uid []byte) []byte {
	return append(uidKeyPrefix, uid...)
}

func (s *Shortner) makeURLKey(rawURL []byte) []byte {
	return append(urlKeyPrefix, rawURL...)
}
