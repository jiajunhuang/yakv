package yakv

import (
	"sync"
)

type Dict struct {
	noCopy
	data sync.Map
}

func (d *Dict) Get(key []byte) interface{} {
	value, exists := d.data.Load(b2s(key))
	if exists {
		return value
	}

	return nil
}

func (d *Dict) Delete(keys ...[]byte) int {
	count := 0

	for _, k := range keys {
		// FIXME not concurrency safe here, too
		key := b2s(k)
		if _, exists := d.data.Load(key); exists {
			count++
			d.data.Delete(key)
		}
	}

	return count
}

func (d *Dict) Exists(key []byte) bool {
	_, exists := d.data.Load(b2s(key))
	return exists
}

func (d *Dict) Set(key []byte, value interface{}) bool {
	// FIXME not concurrency safe
	exists := d.Exists(key)

	d.data.Store(b2s(key), value)

	return exists
}
