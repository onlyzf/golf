package queue

import (
	"sync"
)

// Queue concurrent queue
type Queue interface {
	Empty() bool
	Pop() interface{}
	Push(interface{})
}

type qk struct {
	data []interface{}
	sync.Mutex
}

// NewQueue to new a concurrent queue
func NewQueue() Queue {
	return &qk{data: make([]interface{}, 0, 256)}
}

func (q *qk) Empty() bool {
	return len(q.data) <= 0
}

func (q *qk) Pop() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if len(q.data) > 0 {
		v := q.data[0]
		q.data = q.data[1:]
		return v
	}
	return nil
}

func (q *qk) Push(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.data = append(q.data, v)
}
