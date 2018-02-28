package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// add
func TestSetAdd(t *testing.T) {
	s := Set{}

	assert.Equalf(t, 1, s.Add([]byte("hello")), "should euqal")
	assert.Equalf(t, 0, s.Add([]byte("hello")), "should euqal")
}

func BenchmarkSetAdd(b *testing.B) {
	s := Set{}

	for i := 0; i < b.N; i++ {
		s.Add([]byte("hello"))
	}
}
