package shortner

import (
	"os"
	"testing"

	"github.com/jimeh/ozu.io/storage"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
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

// Setup Suite

var testDbPath = "./goleveldb_test_data"

type ShortnerSuite struct {
	suite.Suite
	shortner *Shortner
	store    storage.Store
}

func (s *ShortnerSuite) SetupTest() {
	store, err := goleveldbstore.New(testDbPath)
	s.Require().NoError(err)

	err = store.Set(goleveldbstore.DefaultSequenceKey, []byte("1000"))
	s.Require().NoError(err)

	s.store = store
	s.shortner = New(store)
}

func (s *ShortnerSuite) TearDownTest() {
	_ = s.store.Close()
	_ = os.RemoveAll(testDbPath)
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
