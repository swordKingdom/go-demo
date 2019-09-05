package stack

import (
	"fmt"
	"go-demo/datastruct/linkedlist"
	"sync"
)

type BaseStack interface {
	Pop() interface{}
	Push(val interface{})
	Top() interface{}
	IsEmpty() bool
	Clear() bool
	Size() int
}

type stack struct {
	stackTop *linkedlist.BNode
	data     *linkedlist.BNode
	Len      int
	lock     *sync.RWMutex
}

func (s *stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.stackTop == nil {
		return nil
	}
	tmp := s.stackTop
	s.stackTop = s.stackTop.Pre
	s.stackTop.Next = nil
	s.Len--
	return tmp.Val
}

func (s *stack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.data == nil {
		node := &linkedlist.BNode{Val: val}
		s.data = node
		s.stackTop = node
	} else {
		s.stackTop.Next = &linkedlist.BNode{Val: val, Pre: s.stackTop}
		s.stackTop = s.stackTop.Next
	}

	s.Len++
}

func (s *stack) Top() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.stackTop.Val
}

func (s *stack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Len == 0
}

func (s *stack) Clear() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = nil
	s.stackTop = nil
	s.Len = 0
	return true
}

func (s *stack) StackPrint() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	tmp := s.data
	fmt.Print("输出栈结构：")
	for tmp != nil {
		fmt.Printf("%v  ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (s *stack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Len
}

func InitStack() *stack {
	return &stack{lock: new(sync.RWMutex)}
}
