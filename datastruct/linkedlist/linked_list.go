package linkedlist

import (
	"sync"
)

//LinkList 链表接口
type linkList interface {
	InsertNode(val interface{}, index int)
	DeleteNode(index int)
	GetNode(index int) *Node
}

//Node 链表的节点
type Node struct {
	Val      interface{}
	NodeType interface{}
	Next     *Node
}

//Clone 链表的节点克隆函数
func (n *Node) Clone() *Node {
	return &Node{Val: n.Val, NodeType: n.NodeType}
}

//Head 链表的链表头函数
type Head struct {
	head *Node
	lock *sync.RWMutex
	len  int
}

//InsertNode 在链表中插入元素
func (h *Head) InsertNode(index int, val interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if index > h.len || index < 0 {
		return
	}
	tmp := h.head
	index--
	for ; index > 0; index-- {
		tmp = tmp.Next
	}

	if tmp.Next == nil {
		tmp.Next = &Node{Val: val}
	} else {
		tmp.Next = &Node{Val: val, Next: tmp.Next}
	}
	h.len++
}

//GetNode 获取链表中的节点
func (h *Head) GetNode(index int) *Node {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if index > h.len || index < 0 {
		return nil
	}
	tmp := h.head
	for ; index > 0; index-- {
		tmp = tmp.Next
	}
	return tmp.Clone()
}

//DeleteNode 删除链表中的节点
func (h *Head) DeleteNode(index int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if index > h.len || index < 0 {
		return
	}
	tmp := h.head
	index--
	for ; index > 0; index-- {
		tmp = tmp.Next
	}
	if tmp.Next != nil {
		if tmp.Next.Next != nil {
			tmp.Next = tmp.Next.Next
		} else {
			tmp.Next = nil
		}
	}
}

//Poll 移除并返问队列头部的元素 如果队列为空，则返回null
func (h *Head) Poll() *Node {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.head == nil {
		return nil
	}
	res := h.head
	h.head = h.head.Next
	h.len--
	return res.Clone()
}

//Peek 返回队列头部的元素 如果队列为空，则返回null
func (h *Head) Peek() *Node {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.head == nil {
		return nil
	}
	return h.head.Clone()
}

func (h *Head) IsSupportSort() bool {
	isSuppoertSort := false
	if h.head == nil {
		return isSuppoertSort
	}
	tKin := h.head.NodeType
	tmp := h.head.Next
	for tmp != nil {
		if tKin != tmp.NodeType {
			isSuppoertSort = false
			break
		}
		tmp = tmp.Next
	}
	return isSuppoertSort
}

//SoredLinkedList 排序函数
//todo:在链表数据中不是同一结构类型的时候不支持排序
func (h *Head) SortedLinkedList(sortFunc func(a *Node, b *Node) int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for p := h.head; p.Next != nil; p = p.Next {
		for q := h.head; q.Next != nil; q = q.Next {
			if sortFunc(p, q) >= 0 {
				q.Val, p.Val = p.Val, q.Val
				q.NodeType, p.NodeType = p.NodeType, q.NodeType
			}
		}
	}
}

func (h *Head) RangeReadOnly(f func(node *Node)) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	tmp := h.head
	for tmp != nil {
		f(tmp.Clone())
		tmp = tmp.Next
	}
}

func (h *Head) Range(f func(node *Node)) {
	h.lock.Lock()
	defer h.lock.Unlock()
	tmp := h.head
	for tmp != nil {
		f(tmp.Clone())
		tmp = tmp.Next
	}
}

//InitLinkedListWithArr 从数组中初始化链表数据结构
func InitLinkedListWithArr(arr ...interface{}) *Head {
	arrLen := len(arr)
	if arrLen == 0 {
		return nil
	}
	head := &Node{Val: arr[0]}
	res := &Head{head: head, len: arrLen, lock: new(sync.RWMutex)}
	for i := 1; i < arrLen; i++ {
		head.Next = &Node{Val: arr[i]}
		head = head.Next
	}
	return res
}
