package shortener

import (
	"errors"
	"strings"
	"testing"

	"github.com/jimeh/ozu.io/shortener/mocks"
	"github.com/jimeh/ozu.io/storage"
	"github.com/stretchr/testify/suite"
)

// Mocks

//go:generate mockery -name Store -dir ../storage

// Suite Setup

type Base58ShortenerSuite struct {
	suite.Suite
	store     *mocks.Store
	shortener *Base58Shortener
}

func (s *Base58ShortenerSuite) SetupTest() {
	s.store = new(mocks.Store)
	s.shortener = NewBase58(s.store)
}

// Tests

func (s *Base58ShortenerSuite) TestShortenExisting() {
	uid := []byte("ig")
	url := []byte("https://google.com/")
	record := storage.Record{UID: uid, URL: url}

	s.store.On("FindByURL", url).Return(&record, nil)

	result, err := s.shortener.Shorten(url)
	s.NoError(err)
	s.Equal(uid, result.UID)
	s.Equal(url, result.URL)
	s.store.AssertExpectations(s.T())
}

func (s *Base58ShortenerSuite) TestShortenNew() {
	uid := []byte("ig")
	url := []byte("https://google.com/")
	record := storage.Record{UID: uid, URL: url}

	s.store.On("FindByURL", url).Return(nil, storage.ErrNotFound)
	s.store.On("NextSequence").Return(1001, nil)
	s.store.On("Create", uid, url).Return(&record, nil)

	result, err := s.shortener.Shorten(url)

	s.NoError(err)
	s.Equal(uid, result.UID)
	s.Equal(url, result.URL)
	s.store.AssertExpectations(s.T())
}

func (s *Base58ShortenerSuite) TestShortenAndNormalizeURL() {
	examples := []struct {
		url        []byte
		normalized []byte
	}{
		{[]byte("google.com"), []byte("http://google.com/")},
		{[]byte("google.com/"), []byte("http://google.com/")},
		{[]byte("http://google.com"), []byte("http://google.com/")},
	}

	for _, e := range examples {
		record := storage.Record{UID: []byte("ig"), URL: e.normalized}
		s.store.On("FindByURL", record.URL).Return(&record, nil)

		result, err := s.shortener.Shorten(e.url)
		s.NoError(err)
		s.Equal(record.UID, result.UID)
		s.Equal(record.URL, result.URL)
		s.store.AssertExpectations(s.T())
	}
}

func (s *Base58ShortenerSuite) TestShortenInvalidURL() {
	examples := []struct {
		url   string
		error string
	}{
		{
			url:   "*$)]+_<?)",
			error: "invalid URL",
		},
		{
			url:   "",
			error: "invalid URL",
		},
		{
			url:   "file:///bin/bash",
			error: "schema 'file://' not allowed",
		},
		{
			url:   "/users/view.php?uid=138495",
			error: "invalid URL",
		},
		{
			url:   "http://long.com/" + strings.Repeat("0", 3000),
			error: "invalid URL",
		},
	}

	for _, e := range examples {
		record, err := s.shortener.Shorten([]byte(e.url))
		s.Nil(record.UID)
		s.Nil(record.URL)
		s.EqualError(err, e.error)
	}
}

func (s *Base58ShortenerSuite) TestShortenStoreError() {
	url := []byte("https://google.com/")
	storeErr := errors.New("leveldb: something wrong")

	s.store.On("FindByURL", url).Return(nil, storeErr)

	result, err := s.shortener.Shorten(url)
	s.Nil(result.UID)
	s.Nil(result.URL)
	s.EqualError(err, storeErr.Error())
	s.store.AssertExpectations(s.T())
}

func (s *Base58ShortenerSuite) TestLookupExisting() {
	uid := []byte("ig")
	url := []byte("https://google.com/")
	record := storage.Record{UID: uid, URL: url}

	s.store.On("FindByUID", uid).Return(&record, nil)

	result, err := s.shortener.Lookup(uid)
	s.NoError(err)
	s.Equal(uid, result.UID)
	s.Equal(url, result.URL)
	s.store.AssertExpectations(s.T())
}

func (s *Base58ShortenerSuite) TestLookupNonExistant() {
	uid := []byte("ig")

	s.store.On("FindByUID", uid).Return(&storage.Record{}, storage.ErrNotFound)

	result, err := s.shortener.Lookup(uid)
	s.EqualError(err, "not found")
	s.Nil(result.UID)
	s.Nil(result.URL)
	s.store.AssertExpectations(s.T())
}

func (s *Base58ShortenerSuite) TestLookupInvalid() {
	uid := []byte("ig\"; drop table haha")

	result, err := s.shortener.Lookup(uid)

	s.EqualError(err, "invalid UID")
	s.Nil(result.UID)
	s.Nil(result.URL)
	s.store.AssertExpectations(s.T())
}

// Run Suite

func TestShortenerSuite(t *testing.T) {
	suite.Run(t, new(Base58ShortenerSuite))
}
