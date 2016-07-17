package goleveldbstore

import (
	"os"
	"testing"

	"github.com/jimeh/ozu.io/storage"
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

type StoreSuite struct {
	suite.Suite
	store *Store
	db    *leveldb.DB
}

func (s *StoreSuite) SetupTest() {
	db, errDB := leveldb.OpenFile(testDbPath, nil)
	s.Require().NoError(errDB)
	store, err := New(db)
	s.Require().NoError(err)
	s.store = store
	s.db = db
}

func (s *StoreSuite) TearDownTest() {
	_ = s.db.Close()
	_ = os.RemoveAll(testDbPath)
}

func (s *StoreSuite) Seed() {
	for _, e := range examples {
		err := s.db.Put(e.key, e.value, nil)
		s.Require().NoError(err)
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
		result, _ := s.db.Get(e.key, nil)
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
		value, _ := s.db.Get(e.key, nil)
		s.Require().Equal(e.value, value)

		value, err := s.store.Get(e.key)
		s.Require().NoError(err)
		s.Require().Equal(value, e.value)

		err = s.store.Delete(e.key)
		s.NoError(err)

		value, err = s.store.Get(e.key)
		s.Nil(value)
		s.EqualError(err, "not found")

		has, _ := s.db.Has(e.key, nil)
		s.Equal(false, has)
		result, _ := s.db.Get(e.key, nil)
		s.Equal([]byte{}, result)
	}
}

func (s *StoreSuite) TestDeleteNonExistant() {
	err := s.store.Delete([]byte("does-not-exist"))
	s.NoError(err)
}

func (s *StoreSuite) TestNextSequenceExisting() {
	err := s.db.Put(DefaultSequenceKey, []byte("5"), nil)
	s.Require().NoError(err)

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

func (s *StoreSuite) TestIncrExisting() {
	key := []byte("my-counter")

	err := s.db.Put(key, []byte("5"), nil)
	s.Require().NoError(err)

	result, err := s.store.Incr(key)
	s.NoError(err)
	s.Equal(6, result)
}

func (s *StoreSuite) TestIncrNonExistant() {
	for i := 1; i < 10; i++ {
		result, err := s.store.Incr([]byte("counter"))

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
	db, _ := leveldb.OpenFile(testDbPath, nil)
	store, _ := New(db)

	key := []byte("hello")
	value := []byte("world")
	_ = store.Set(key, value)

	for n := 0; n < b.N; n++ {
		_, _ = store.Get(key)
	}

	_ = os.RemoveAll(testDbPath)
}

func BenchmarkSet(b *testing.B) {
	db, _ := leveldb.OpenFile(testDbPath, nil)
	store, _ := New(db)

	key := []byte("hello")
	value := []byte("world")

	for n := 0; n < b.N; n++ {
		_ = store.Set(append(key, string(n)...), value)
	}

	_ = os.RemoveAll(testDbPath)
}

func BenchmarkNextSequence(b *testing.B) {
	db, _ := leveldb.OpenFile(testDbPath, nil)
	store, _ := New(db)

	for n := 0; n < b.N; n++ {
		_, _ = store.NextSequence()
	}

	_ = os.RemoveAll(testDbPath)
}

func BenchmarkIncr(b *testing.B) {
	db, _ := leveldb.OpenFile(testDbPath, nil)
	store, _ := New(db)

	key := []byte("incr-benchmark-counter")
	for n := 0; n < b.N; n++ {
		_, _ = store.Incr(key)
	}

	_ = os.RemoveAll(testDbPath)
}
