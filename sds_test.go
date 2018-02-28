package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// append
func TestSDSAppend(t *testing.T) {
	s := SDS{}
	length := s.Append([]byte("hello"))

	assert.Equal(t, 5, length, "length of sds should be 5")
}

func BenchmarkSDSAppend(b *testing.B) {
	hello := []byte("hello")
	world := []byte("world")

	for i := 0; i < b.N; i++ {
		s := SDS{}
		s.Append(hello)
		s.Append(world)
		s.Append(hello)
	}
}

// len
func TestSDSLen(t *testing.T) {
	s := SDS{}

	s.Append([]byte("hello"))

	assert.Equal(t, 5, s.Len(), "length of sds should be 5")
}

func BenchmarkSDSLen(b *testing.B) {
	hello := []byte("hello")
	s := SDS{}
	s.Append(hello)

	for i := 0; i < b.N; i++ {
		s.Len()
	}
}

// getrange
func TestSDSGetRange(t *testing.T) {
	s := SDS{}
	s.Append([]byte("hello"))

	type condition struct {
		left     int
		right    int
		expected []byte
	}
	conditions := []condition{
		condition{-1, -3, []byte("")},
		condition{0, 4, []byte("hello")},
		condition{0, 5, []byte("hello")},
		condition{-1, 5, []byte("o")},
		condition{-10, 5, []byte("hello")},
	}

	for _, c := range conditions {
		result := s.GetRange(c.left, c.right)
		assert.Equalf(
			t, c.expected, result, "s.GetRange(%d, %d) should get %s but got: %s", c.left, c.right, c.expected, result,
		)
	}
}

func BenchmarkSDSGetRange(b *testing.B) {
	hello := []byte("hello")
	s := SDS{}
	s.Append(hello)

	for i := 0; i < b.N; i++ {
		s.GetRange(-10, 10)
	}
}
