package yakv

import (
	"sync"
)

type Set struct {
	data sync.Map
}

// Add : 0 means exist already so nothing happend
func (s *Set) Add(keys ...[]byte) int {
	count := 0

	for i := range keys {
		_, loaded := s.data.LoadOrStore(b2s(keys[i]), Nil)
		if !loaded {
			count++
		}
	}

	return count
}
