package shortner

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jimeh/ozu.io/shortner/mocks"
	"github.com/stretchr/testify/suite"
)

// Mocks

//go:generate mockery -name Store -dir .. -recursive

// Suite Setup

type ShortnerSuite struct {
	suite.Suite
	store       *mocks.Store
	shortner    *Shortner
	errNotFound error
}

func (s *ShortnerSuite) SetupTest() {
	s.store = new(mocks.Store)
	s.shortner = New(s.store)
	s.errNotFound = errors.New("not found")
}

// Tests

func (s *ShortnerSuite) TestShortenExisting() {
	rawURL := []byte("http://google.com/")
	uid := []byte("ig")
	urlSHA := fmt.Sprintf("%x", sha1.Sum(rawURL))

	s.store.On("Get", append([]byte("url:"), urlSHA...)).Return(uid, nil)

	resultUID, resultURL, err := s.shortner.Shorten(rawURL)
	s.NoError(err)
	s.Equal(uid, resultUID)
	s.Equal(rawURL, resultURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortnerSuite) TestShortenNew() {
	rawURL := []byte("https://google.com")
	url := []byte("https://google.com/")
	uid := []byte("ig")
	urlKey := append([]byte("url:"), fmt.Sprintf("%x", sha1.Sum(url))...)

	s.store.On("Get", urlKey).Return(nil, s.errNotFound)
	s.store.On("NextSequence").Return(1001, nil)
	s.store.On("Set", urlKey, uid).Return(nil)
	s.store.On("Set", append([]byte("uid:"), uid...), url).Return(nil)

	rUID, rURL, err := s.shortner.Shorten(rawURL)

	s.NoError(err)
	s.Equal(uid, rUID)
	s.Equal(url, rURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortnerSuite) TestShortenInvalidURL() {
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
		rUID, rURL, err := s.shortner.Shorten([]byte(e.url))
		s.Nil(rUID)
		s.Nil(rURL)
		s.EqualError(err, e.error)
	}
}

func (s *ShortnerSuite) TestShortenStoreError() {
	url := []byte("https://google.com/")
	storeErr := errors.New("leveldb: something wrong")
	urlKey := append([]byte("url:"), fmt.Sprintf("%x", sha1.Sum(url))...)

	s.store.On("Get", urlKey).Return(nil, storeErr)

	rUID, rURL, err := s.shortner.Shorten(url)
	s.Nil(rUID)
	s.Nil(rURL)
	s.EqualError(err, storeErr.Error())
}

func (s *ShortnerSuite) TestLookupExisting() {
	url := []byte("https://google.com/")
	uid := []byte("ig")

	s.store.On("Get", append([]byte("uid:"), uid...)).Return(url, nil)

	rURL, err := s.shortner.Lookup(uid)

	s.NoError(err)
	s.Equal(url, rURL)
	s.store.AssertExpectations(s.T())
}

func (s *ShortnerSuite) TestLookupNonExistant() {
	uid := []byte("ig")

	s.store.On("Get", append([]byte("uid:"), uid...)).Return(nil, s.errNotFound)

	rURL, err := s.shortner.Lookup(uid)

	s.EqualError(err, "not found")
	s.Nil(rURL)
	s.store.AssertExpectations(s.T())
}

// Run Suite

func TestShortnerSuite(t *testing.T) {
	suite.Run(t, new(ShortnerSuite))
}
