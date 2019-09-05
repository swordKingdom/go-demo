package queue

import (
	"go-demo/datastruct/linkedlist"
	"sync"
)

type BaseQueue interface {
	InQueue(a interface{})
	OutQueue() interface{}
	IsEmpty() bool
	Clear() bool
	Size() int
	IsFull() bool
	Traverse(func(d interface{}))
}

type Queue struct {
	data *linkedlist.BNode
	len  int
	size int
	lock *sync.RWMutex
}

func (q *Queue) InQueue(a interface{}) {
	panic("implement me")
}

func (q *Queue) OutQueue() interface{} {
	panic("implement me")
}

func (q *Queue) IsEmpty() bool {
	panic("implement me")
}

func (q *Queue) Clear() bool {
	panic("implement me")
}

func (q *Queue) Size() int {
	panic("implement me")
}

func (q *Queue) IsFull() bool {
	return q.len == q.size
}

func (q *Queue) Traverse(func(d interface{})) {
	panic("implement me")
}
