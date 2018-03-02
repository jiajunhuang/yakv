package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	l := List{}

	assert.Equal(t, 0, l.Len(), "should equal")

	l.Lpush(1)
	assert.Equal(t, 1, l.Len(), "should equal")
}

func TestLpush(t *testing.T) {
	l := List{}

	assert.Equal(t, 0, l.Len(), "should equal")

	l.Lpush(1)
	assert.Equal(t, 1, l.Len(), "should equal")
	l.Lpush(1)
	assert.Equal(t, 2, l.Len(), "should equal")
}
