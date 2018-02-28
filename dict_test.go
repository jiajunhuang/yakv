package yakv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	value1 = "value1"
	value2 = "value2"
)

// get & set
func TestDictGet(t *testing.T) {
	d := Dict{}

	d.Set(key1, value1)

	assert.Equal(t, value1, d.Get(key1), "should equal")
}

// delete
func TestDictDelete(t *testing.T) {
	d := Dict{}

	d.Set(key1, value1)
	assert.Equal(t, value1, d.Get(key1), "should equal")

	d.Delete(key1)
	assert.Equal(t, nil, d.Get(key1), "should equal")
}

// exists
func TestDictExists(t *testing.T) {
	d := Dict{}

	d.Set(key1, value1)
	assert.Equal(t, true, d.Exists(key1), "should equal")
	assert.Equal(t, false, d.Exists(key2), "should equal")
}
