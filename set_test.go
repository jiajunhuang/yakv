package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key1 = []byte("hello")
	key2 = []byte("world")
)

// add
func TestSetAdd(t *testing.T) {
	s := Set{}

	assert.Equalf(t, 1, s.Add(key1), "should euqal")
	assert.Equalf(t, 0, s.Add(key1), "should euqal")
}

func BenchmarkSetAdd(b *testing.B) {
	s := Set{}

	for i := 0; i < b.N; i++ {
		s.Add(key1)
	}
}

// ismember
func TestSetIsMember(t *testing.T) {
	s := Set{}

	s.Add(key1)

	assert.Equal(t, true, s.IsMember(key1), "should exist")
	assert.Equal(t, false, s.IsMember(key2), "should not exist")
}

func BenchmarkSetIsMember(b *testing.B) {
	s := Set{}

	for i := 0; i < b.N; i++ {
		s.IsMember(key1)
	}
}

// remove
func TestSetRemove(t *testing.T) {
	s := Set{}

	s.Add(key1)
	assert.Equal(t, true, s.IsMember(key1), "should exist")

	s.Remove(key1)
	assert.Equal(t, false, s.IsMember(key1), "should not exist")
}

func BenchmarkSetRemove(b *testing.B) {
	s := Set{}
	s.Add(key1)

	for i := 0; i < b.N; i++ {
		s.Remove(key1)
	}
}

// count
func TestSetCount(t *testing.T) {
	s := Set{}

	assert.Equal(t, 0, s.Count(), "should equal")

	s.Add(key1)
	assert.Equal(t, 1, s.Count(), "should equal")

	s.Add(key1)
	assert.Equal(t, 1, s.Count(), "should equal")

	s.Add(key2)
	assert.Equal(t, 2, s.Count(), "should equal")
}

func BenchmarkSetCount(b *testing.B) {
	s := Set{}
	s.Add(key1)
	s.Add(key2)

	for i := 0; i < b.N; i++ {
		s.Count()
	}
}
