package shortener

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jimeh/ozu.io/shortener/mocks"
	"github.com/stretchr/testify/suite"
)

// Mocks

//go:generate mockery -name Store -dir .. -recursive

// Suite Setup

type ShortenerSuite struct {
	suite.Suite
	store       *mocks.Store
	shortener   *Shortener
	errNotFound error
}

func (s *ShortenerSuite) SetupTest() {
	s.store = new(mocks.Store)
	s.shortener = New(s.store)
	s.errNotFound = errors.New("not found")
}

// Tests

func (s *ShortenerSuite) TestShortenExisting() {
	rawURL := []byte("http://google.com/")
	uid := []byte("ig")
	urlSHA := fmt.Sprintf("%x", sha1.Sum(rawURL))

	s.store.On("Get", append([]byte("url:"), urlSHA...)).Return(uid, nil)

	resultUID, resultURL, err := s.shortener.Shorten(rawURL)
	s.NoError(err)
	s.Equal(uid, resultUID)
	s.Equal(rawURL, resultURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortenerSuite) TestShortenNew() {
	rawURL := []byte("https://google.com")
	url := []byte("https://google.com/")
	uid := []byte("ig")
	urlKey := append([]byte("url:"), fmt.Sprintf("%x", sha1.Sum(url))...)

	s.store.On("Get", urlKey).Return(nil, s.errNotFound)
	s.store.On("NextSequence").Return(1001, nil)
	s.store.On("Set", urlKey, uid).Return(nil)
	s.store.On("Set", append([]byte("uid:"), uid...), url).Return(nil)

	rUID, rURL, err := s.shortener.Shorten(rawURL)

	s.NoError(err)
	s.Equal(uid, rUID)
	s.Equal(url, rURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortenerSuite) TestShortenInvalidURL() {
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
		rUID, rURL, err := s.shortener.Shorten([]byte(e.url))
		s.Nil(rUID)
		s.Nil(rURL)
		s.EqualError(err, e.error)
	}
}

func (s *ShortenerSuite) TestShortenStoreError() {
	url := []byte("https://google.com/")
	storeErr := errors.New("leveldb: something wrong")
	urlKey := append([]byte("url:"), fmt.Sprintf("%x", sha1.Sum(url))...)

	s.store.On("Get", urlKey).Return(nil, storeErr)

	rUID, rURL, err := s.shortener.Shorten(url)
	s.Nil(rUID)
	s.Nil(rURL)
	s.EqualError(err, storeErr.Error())
}

func (s *ShortenerSuite) TestLookupExisting() {
	url := []byte("https://google.com/")
	uid := []byte("ig")

	s.store.On("Get", append([]byte("uid:"), uid...)).Return(url, nil)

	rURL, err := s.shortener.Lookup(uid)

	s.NoError(err)
	s.Equal(url, rURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortenerSuite) TestLookupNonExistant() {
	uid := []byte("ig")

	s.store.On("Get", append([]byte("uid:"), uid...)).Return(nil, s.errNotFound)

	rURL, err := s.shortener.Lookup(uid)

	s.EqualError(err, "not found")
	s.Nil(rURL)
	s.store.AssertExpectations(s.T())
}

// Run Suite

func TestShortenerSuite(t *testing.T) {
	suite.Run(t, new(ShortenerSuite))
}
