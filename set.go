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

// IsMember implements SISMEMBER
func (s *Set) IsMember(key []byte) bool {
	_, ok := s.data.Load(b2s(key))
	return ok
}

// Remove return numbers removed
func (s *Set) Remove(keys ...[]byte) int {
	count := 0

	// FIXME: as we're using sync.Map, it does not provide such operation(delete and count)
	// so here, the operation `Remove` is not concurrency safe
	// TODO: maybe change this implementation in sync.Map.Range
	for i := range keys {
		key := b2s(keys[i])
		_, ok := s.data.Load(key)
		if ok {
			count++
		}
		s.data.Delete(key)
	}

	return count
}

// Count return count of elements in sync.Map
func (s *Set) Count() int {
	length := 0

	s.data.Range(func(k, v interface{}) bool {
		length++

		return true
	})

	return length
}
