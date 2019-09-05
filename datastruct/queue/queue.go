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
	front *linkedlist.BNode
	rear  *linkedlist.BNode
	data  *linkedlist.BNode
	len   int
	size  int
	lock  *sync.RWMutex
}

func (q *Queue) InQueue(a interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size < q.len+1 {
		return
	}

	if q.data == nil {
		node := &linkedlist.BNode{Val: a}
		q.data = node
		q.front = node
		q.rear = node
	} else {
		node := &linkedlist.BNode{Val: a, Pre: q.rear}
		q.rear.Next = node
		q.rear = q.rear.Next
	}
	q.len++
}

func (q *Queue) OutQueue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.len == 0 {
		return nil
	}
	res := q.data.Val
	q.data = q.data.Next
	q.front = q.data
	return res
}

func (q *Queue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.len == 0
}

func (q *Queue) Clear() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data = nil
	q.front = nil
	q.rear = nil
	q.len = 0
	return true
}

func (q *Queue) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.size
}

func (q *Queue) IsFull() bool {
	return q.len == q.size
}

func (q *Queue) Traverse(f func(d interface{})) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	tmp := q.data
	for tmp != nil {
		f(tmp.Val)
		tmp = tmp.Next
	}
}

func InitQueue(size int) *Queue {
	return &Queue{size: size, lock: new(sync.RWMutex)}
}
