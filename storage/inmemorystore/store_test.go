package inmemorystore

import (
	"testing"

	"github.com/jimeh/ozu.io/storage"
	"github.com/stretchr/testify/suite"
)

// Setup Suite

var examples = []storage.Record{
	storage.Record{UID: []byte("Kb8X"), URL: []byte("https://google.com/")},
	storage.Record{UID: []byte("h3mz"), URL: []byte("https://github.com/")},
	storage.Record{UID: []byte("3qxs"), URL: []byte("https://twitter.com/")},
}

type StoreSuite struct {
	suite.Suite
	store *Store
}

func (s *StoreSuite) SetupTest() {
	store, err := New()
	s.Require().NoError(err)
	s.store = store
}

func (s *StoreSuite) TearDownTest() {
	_ = s.store.Close()
}

func (s *StoreSuite) Seed() {
	for _, e := range examples {
		s.store.UIDMap[string(e.UID)] = e.URL
		s.store.URLMap[string(e.URL)] = e.UID
	}
}

// Tests

func (s *StoreSuite) TestStoreInterface() {
	s.Implements(new(storage.Store), new(Store))
}

func (s *StoreSuite) TestCreate() {
	for _, e := range examples {
		record, err := s.store.Create(e.UID, e.URL)
		s.NoError(err)
		s.Equal(e.UID, record.UID)
		s.Equal(e.URL, record.URL)
	}

	for _, e := range examples {
		recordURL, _ := s.store.UIDMap[string(e.UID)]
		s.Equal(e.URL, recordURL)
		recordUID, _ := s.store.URLMap[string(e.URL)]
		s.Equal(e.UID, recordUID)
	}
}

func (s *StoreSuite) TestFindExistingByUID() {
	s.Seed()

	for _, e := range examples {
		record, err := s.store.FindByUID(e.UID)
		s.NoError(err)
		s.Equal(e.UID, record.UID)
		s.Equal(e.URL, record.URL)
	}
}

func (s *StoreSuite) TestFindNonExistantByUID() {
	record, err := s.store.FindByUID([]byte("does-not-exist"))
	s.Nil(record.UID)
	s.Nil(record.URL)
	s.EqualError(err, "not found")
}

func (s *StoreSuite) TestFindExistingByURL() {
	s.Seed()

	for _, e := range examples {
		record, err := s.store.FindByURL(e.URL)
		s.NoError(err)
		s.Equal(e.UID, record.UID)
		s.Equal(e.URL, record.URL)
	}
}

func (s *StoreSuite) TestFindNonExistantByURL() {
	record, err := s.store.FindByURL([]byte("http://nope.com/"))
	s.Nil(record.UID)
	s.Nil(record.URL)
	s.EqualError(err, "not found")
}

func (s *StoreSuite) TestDeleteExistingByUID() {
	s.Seed()

	for _, e := range examples {
		record, err := s.store.DeleteByUID(e.UID)
		s.NoError(err)
		s.Equal(record.UID, e.UID)
		s.Equal(record.URL, e.URL)

		record, err = s.store.FindByUID(e.UID)
		s.Nil(record.UID)
		s.Nil(record.URL)
		s.EqualError(err, "not found")

		record, err = s.store.FindByURL(e.URL)
		s.Nil(record.UID)
		s.Nil(record.URL)
		s.EqualError(err, "not found")
	}
}

func (s *StoreSuite) TestDeleteNonExistantByUID() {
	record, err := s.store.DeleteByUID([]byte("nope"))
	s.Nil(record.UID)
	s.Nil(record.URL)
	s.EqualError(err, "not found")
}

func (s *StoreSuite) TestDeleteExistingByURL() {
	s.Seed()

	for _, e := range examples {
		record, err := s.store.DeleteByURL(e.URL)
		s.NoError(err)
		s.Equal(record.UID, e.UID)
		s.Equal(record.URL, e.URL)

		record, err = s.store.FindByUID(e.UID)
		s.Nil(record.UID)
		s.Nil(record.URL)
		s.EqualError(err, "not found")

		record, err = s.store.FindByURL(e.URL)
		s.Nil(record.UID)
		s.Nil(record.URL)
		s.EqualError(err, "not found")
	}
}

func (s *StoreSuite) TestDeleteNonExistantByURL() {
	record, err := s.store.DeleteByURL([]byte("http://nope/"))
	s.Nil(record.UID)
	s.Nil(record.URL)
	s.EqualError(err, "not found")
}

func (s *StoreSuite) TestNextSequenceExisting() {
	s.store.Sequence = 5

	result, err := s.store.NextSequence()
	s.NoError(err)
	s.Equal(6, result)
}

func (s *StoreSuite) TestNextSequenceNonExistant() {
	for i := 1; i < 10; i++ {
		result, err := s.store.NextSequence()

		s.NoError(err)
		s.Equal(i, result)
	}
}

// Run Suite

func TestStoreSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

// Benchmarks

func BenchmarkCreate(b *testing.B) {
	store, _ := New()

	uid := []byte("Kb8X")
	url := []byte("https://google.com/")

	for n := 0; n < b.N; n++ {
		store.Create(append(uid, string(n)...), url)
	}

	store.Close()
}

func BenchmarkFindByUID(b *testing.B) {
	store, _ := New()

	uid := []byte("Kb8X")
	url := []byte("https://google.com/")
	store.Create(uid, url)

	for n := 0; n < b.N; n++ {
		store.FindByUID(uid)
	}

	store.Close()
}

func BenchmarkFindByURL(b *testing.B) {
	store, _ := New()

	uid := []byte("Kb8X")
	url := []byte("https://google.com/")
	store.Create(uid, url)

	for n := 0; n < b.N; n++ {
		store.FindByURL(url)
	}

	store.Close()
}

func BenchmarkNextSequence(b *testing.B) {
	store, _ := New()

	for n := 0; n < b.N; n++ {
		store.NextSequence()
	}

	store.Close()
}
