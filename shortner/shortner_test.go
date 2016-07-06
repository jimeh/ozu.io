package shortner

import (
	"errors"
	"strconv"
	"sync"
	"testing"

	"github.com/jimeh/ozu.io/storage"
	"github.com/stretchr/testify/suite"
)

// Test Cases

var shortenExamples = []struct {
	uid        string
	url        string
	normalized string
}{
	{uid: "ig", url: "google.com", normalized: "http://google.com/"},
	{uid: "ih", url: "https://google.com", normalized: "https://google.com/"},
	{uid: "ig", url: "http://google.com", normalized: "http://google.com/"},
	{uid: "ih", url: "https://google.com/"},
	{uid: "ig", url: "google.com/", normalized: "http://google.com/"},
	{uid: "ii", url: "https://github.com/"},
	{uid: "ij", url: "https://gist.github.com/"},
}

// Mocks

func NewMockStore() *MockStore {
	return &MockStore{
		Data:     map[string][]byte{},
		Sequence: 1000,
	}
}

type MockStore struct {
	sync.RWMutex
	Data     map[string][]byte
	Sequence int
}

func (s *MockStore) Close() error {
	return nil
}

func (s *MockStore) Get(key []byte) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	value := s.Data[string(key)]
	if value == nil {
		return nil, errors.New("not found")
	}
	return value, nil
}

func (s *MockStore) Set(key []byte, value []byte) error {
	s.Lock()
	defer s.Unlock()
	s.Data[string(key)] = value
	return nil
}

func (s *MockStore) Delete(key []byte) error {
	s.Lock()
	defer s.Unlock()
	delete(s.Data, string(key))
	return nil
}

func (s *MockStore) NextSequence() (int, error) {
	s.Lock()
	defer s.Unlock()
	s.Sequence++
	return s.Sequence, nil
}

// Setup Suite

type ShortnerSuite struct {
	suite.Suite
	shortner *Shortner
	store    storage.Store
}

func (s *ShortnerSuite) SetupTest() {
	s.store = NewMockStore()
	s.shortner = New(s.store)
}

func (s *ShortnerSuite) Seed() {
	r := s.Require()

	for _, e := range shortenExamples {
		uid, url, err := s.shortner.Shorten([]byte(e.url))
		r.Equal([]byte(e.uid), uid)
		if e.normalized != "" {
			r.Equal([]byte(e.normalized), url)
		} else {
			r.Equal([]byte(e.url), url)
		}
		r.NoError(err)
	}
}

// Tests

func (s *ShortnerSuite) TestShorten() {
	for _, e := range shortenExamples {
		uid, url, err := s.shortner.Shorten([]byte(e.url))
		s.Equal(nil, err)
		s.Equal([]byte(e.uid), uid)
		if e.normalized != "" {
			s.Equal([]byte(e.normalized), url)
		} else {
			s.Equal([]byte(e.url), url)
		}
	}
}

func (s *ShortnerSuite) TestLookup() {
	s.Seed()

	for _, e := range shortenExamples {
		url, err := s.shortner.Lookup([]byte(e.uid))
		s.NoError(err)

		if e.normalized != "" {
			s.Equal([]byte(e.normalized), url)
		} else {
			s.Equal([]byte(e.url), url)
		}
	}
}

// Run Suite

func TestShortnerSuite(t *testing.T) {
	suite.Run(t, new(ShortnerSuite))
}

// Benchmarks

func BenchmarkShorten(b *testing.B) {
	shortner := New(NewMockStore())
	rawURL := []byte("https://google.com/")

	for n := 0; n < b.N; n++ {
		_, _, _ = shortner.Shorten(append(rawURL, strconv.Itoa(n)...))
	}
}

func BenchmarkLookup(b *testing.B) {
	shortner := New(NewMockStore())
	rawURL := []byte("https://google.com/")
	uid, _, _ := shortner.Shorten(rawURL)

	for n := 0; n < b.N; n++ {
		_, _ = shortner.Lookup(uid)
	}
}
