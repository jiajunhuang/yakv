package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// b2s
func TestB2S(t *testing.T) {
	assert.Equalf(t, "hello", b2s([]byte("hello")), "should equal")
}

func BenchmarkB2S(b *testing.B) {
	bs := []byte("hello")

	for i := 0; i < b.N; i++ {
		b2s(bs)
	}
}

// s2b
func TestS2B(t *testing.T) {
	assert.Equalf(t, []byte("hello"), s2b("hello"), "should equal")
}

func BenchmarkS2B(b *testing.B) {
	s := "hello"

	for i := 0; i < b.N; i++ {
		s2b(s)
	}
}
