package yakv

import (
	"testing"
)

func TestLock(t *testing.T) {
	type S struct {
		noCopy
	}

	s := S{}
	s.Lock()
}
