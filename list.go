package yakv

import (
	"sync"
)

// List is a doubly linked list implementation
type List struct {
	l      sync.Mutex
	length int
	head   *node
	tail   *node
}

type node struct {
	v    interface{}
	prev *node
	next *node
}

func (l *List) Len() int {
	l.l.Lock()
	length := l.length
	l.l.Unlock()

	return length
}

func (l *List) Lpush(vs ...interface{}) int {
	l.l.Lock()
	length := l.length

	for _, v := range vs {
		if l.head == nil {
			l.head = &node{}
			l.tail = l.head
			l.head.v = v
		} else {
			n := &node{v: v, prev: nil, next: l.head}
			l.head.prev = n
			l.head = n
		}
		length++
	}
	l.length = length
	l.l.Unlock()

	return length
}
