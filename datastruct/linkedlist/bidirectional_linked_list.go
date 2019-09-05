package linkedlist

import "sync"

type BNode struct {
	Val  interface{}
	Pre  *BNode
	Next *BNode
}

type BLinkedListHead struct {
	head *BNode
	len  int
	lock *sync.RWMutex
}

func (b *BNode) Clone() *BNode {
	return &BNode{Val: b.Val}
}

func (b *BLinkedListHead) InsertNode(val interface{}, index int) {

}

func (b *BLinkedListHead) DeleteNode(index int) {
	panic("implement me")
}

func (b *BLinkedListHead) GetNode(index int) *Node {
	panic("implement me")
}

func InitBLinkedList(arr []interface{}) *BLinkedListHead {
	arrLen := len(arr)
	if arrLen == 0 {
		return nil
	}
	head := &BNode{Val: arr[0]}
	res := &BLinkedListHead{head: head, len: arrLen, lock: new(sync.RWMutex)}
	for i := 1; i < arrLen; i++ {
		node := &BNode{Val: arr[i]}
		head.Next = node
		head.Next.Pre = node
		head = head.Next
	}
	return res
}
