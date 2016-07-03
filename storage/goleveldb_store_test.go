package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/syndtr/goleveldb/leveldb"
)

// Setup Suite

var testDbPath = "./goleveldb_test_data"

var examples = []struct {
	key   []byte
	value []byte
}{
	{key: []byte("hello"), value: []byte("world")},
	{key: []byte("foo"), value: []byte("bar")},
	{key: []byte("wtf"), value: []byte("dude")},
}

type GoleveldbStoreSuite struct {
	suite.Suite
	store Store
	db    *leveldb.DB
}

func (s *GoleveldbStoreSuite) Seed() {
	for _, e := range examples {
		err := s.db.Put(e.key, e.value, nil)
		s.Require().Nil(err)
	}
}

func (s *GoleveldbStoreSuite) SetupTest() {
	store, err := NewGoleveldbStore(testDbPath)
	s.Nil(err)
	s.store = &store
	s.db = store.DB
}

func (s *GoleveldbStoreSuite) TearDownTest() {
	s.store.Close()
	os.RemoveAll(testDbPath)
}

// Tests

func (s *GoleveldbStoreSuite) TestSet() {
	for _, e := range examples {
		err := s.store.Set(e.key, e.value)
		s.Nil(err)
	}

	for _, e := range examples {
		result, _ := s.db.Get(e.key, nil)
		s.Equal(e.value, result)
	}
}

func (s *GoleveldbStoreSuite) TestGetExisting() {
	s.Seed()

	for _, e := range examples {
		result, err := s.store.Get(e.key)
		s.Nil(err)
		s.Equal(e.value, result)
	}
}

func (s *GoleveldbStoreSuite) TestGetNonExistant() {
	result, err := s.store.Get([]byte("does-not-exist"))
	s.Nil(err)
	s.Nil(result)
}

func (s *GoleveldbStoreSuite) TestDeleteExisting() {
	s.Seed()

	for _, e := range examples {
		value, _ := s.db.Get(e.key, nil)
		s.Require().Equal(e.value, value)

		err := s.store.Delete(e.key)
		s.Nil(err)

		has, _ := s.db.Has(e.key, nil)
		s.Equal(false, has)
		result, _ := s.db.Get(e.key, nil)
		s.Equal([]byte{}, result)
	}
}

func (s *GoleveldbStoreSuite) TestDeleteNonExistant() {
	err := s.store.Delete([]byte("does-not-exist"))
	s.Nil(err)
}

func (s *GoleveldbStoreSuite) TestIncrExisting() {
	key := []byte("my-counter")

	err := s.db.Put(key, []byte("5"), nil)
	s.Require().Nil(err)

	result, err := s.store.Incr(key)
	s.Nil(err)
	s.Equal(6, result)
}

func (s *GoleveldbStoreSuite) TestIncrNonExistant() {
	for i := 1; i < 10; i++ {
		result, err := s.store.Incr([]byte("counter"))

		s.Nil(err)
		s.Equal(i, result)
	}
}

// Run Suite

func TestGoleveldbStoreSuite(t *testing.T) {
	suite.Run(t, new(GoleveldbStoreSuite))
}

// Benchmarks

func BenchmarkGet(b *testing.B) {
	key := []byte("hello")
	value := []byte("world")
	store, _ := NewGoleveldbStore(testDbPath)
	_ = store.Set(key, value)

	for n := 0; n < b.N; n++ {
		_, _ = store.Get(key)
	}

	_ = store.Close()
	_ = os.RemoveAll(testDbPath)
}

func BenchmarkSet(b *testing.B) {
	key := []byte("hello")
	value := []byte("world")
	store, _ := NewGoleveldbStore(testDbPath)
	_ = store.Set(key, value)

	for n := 0; n < b.N; n++ {
		_ = store.Set(append(key, string(n)...), append(value, string(n)...))
	}

	_ = store.Close()
	_ = os.RemoveAll(testDbPath)
}

func BenchmarkIncr(b *testing.B) {
	key := []byte("incr-benchmark-counter")
	store, _ := NewGoleveldbStore(testDbPath)

	for n := 0; n < b.N; n++ {
		_, _ = store.Incr(key)
	}

	_ = store.Close()
	_ = os.RemoveAll(testDbPath)
}
