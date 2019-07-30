package queue

import (
	"sync"
)

type node struct {
	value interface{}
	next  *node
}

func New() *Queue {
	return &Queue{}
}

type Queue struct {
	head *node
	end  *node
	lock sync.Mutex
}

func (q *Queue) Push(value interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &node{value, nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	return
}

func (q *Queue) Pop() (interface{}, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil, false
	}

	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}

	return value, true
}
