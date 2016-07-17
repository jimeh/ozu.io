package inmemorystore

import (
	"testing"

	"github.com/jimeh/ozu.io/storage"
	"github.com/stretchr/testify/suite"
)

// Setup Suite

var examples = []struct {
	key   []byte
	value []byte
}{
	{key: []byte("hello"), value: []byte("world")},
	{key: []byte("foo"), value: []byte("bar")},
	{key: []byte("wtf"), value: []byte("dude")},
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

func (s *StoreSuite) Seed() {
	for _, e := range examples {
		s.store.Data[string(e.key)] = e.value
	}
}

// Tests

func (s *StoreSuite) TestStoreInterface() {
	s.Implements(new(storage.Store), new(Store))
}

func (s *StoreSuite) TestSet() {
	for _, e := range examples {
		err := s.store.Set(e.key, e.value)
		s.NoError(err)
	}

	for _, e := range examples {
		result, _ := s.store.Data[string(e.key)]
		s.Equal(e.value, result)
	}
}

func (s *StoreSuite) TestGetExisting() {
	s.Seed()

	for _, e := range examples {
		result, err := s.store.Get(e.key)
		s.NoError(err)
		s.Equal(e.value, result)
	}
}

func (s *StoreSuite) TestGetNonExistant() {
	result, err := s.store.Get([]byte("does-not-exist"))
	s.Nil(result)
	s.EqualError(err, "not found")
}

func (s *StoreSuite) TestDeleteExisting() {
	s.Seed()

	for _, e := range examples {
		value := s.store.Data[string(e.key)]
		s.Require().Equal(e.value, value)

		value, err := s.store.Get(e.key)
		s.Require().NoError(err)
		s.Require().Equal(value, e.value)

		err = s.store.Delete(e.key)
		s.NoError(err)

		value, err = s.store.Get(e.key)
		s.Nil(value)
		s.EqualError(err, "not found")

		_, has := s.store.Data[string(e.key)]
		s.Equal(false, has)
	}
}

func (s *StoreSuite) TestDeleteNonExistant() {
	err := s.store.Delete([]byte("does-not-exist"))
	s.NoError(err)
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

func BenchmarkGet(b *testing.B) {
	store, _ := New()

	key := []byte("hello")
	value := []byte("world")
	_ = store.Set(key, value)

	for n := 0; n < b.N; n++ {
		_, _ = store.Get(key)
	}
}

func BenchmarkSet(b *testing.B) {
	store, _ := New()

	key := []byte("hello")
	value := []byte("world")

	for n := 0; n < b.N; n++ {
		_ = store.Set(append(key, string(n)...), value)
	}
}

func BenchmarkNextSequence(b *testing.B) {
	store, _ := New()

	for n := 0; n < b.N; n++ {
		_, _ = store.NextSequence()
	}
}
