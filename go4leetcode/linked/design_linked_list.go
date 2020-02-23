package linked

/*
题目：707. Design Linked List (设计链表)
*/

type MyLinkedList struct {
	Head *ListNode
	Len  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Len {
		return -1
	}
	tmpPoint := this.Head
	for i := 0; i < index; i++ {
		tmpPoint = tmpPoint.Next
	}
	return tmpPoint.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmpNode := &ListNode{Val: val}
	defer func() {
		this.Len++
	}()
	if this.Len == 0 {
		this.Head = tmpNode
		return
	}
	tmpNode.Next = this.Head
	this.Head = tmpNode
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmpNode := &ListNode{Val: val}
	defer func() {
		this.Len++
	}()
	if this.Len == 0 {
		this.Head = tmpNode
		return
	}
	tmpPoint := this.Head
	for tmpPoint.Next != nil {
		tmpPoint = tmpPoint.Next
	}
	tmpPoint.Next = tmpNode
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Len < index {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	tmpPoint := this.Head
	for i := 0; i < index-1; i++ {
		tmpPoint = tmpPoint.Next
	}
	tmpPoint.Next = &ListNode{Val: val, Next: tmpPoint.Next}
	this.Len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Len || index < 0 {
		return
	}
	defer func() {
		this.Len--
	}()
	if index == 0 {
		this.Head = this.Head.Next
		return
	}
	tmpPoint := this.Head
	for i := 0; i < index-1; i++ {
		tmpPoint = tmpPoint.Next
	}
	tmpPoint.Next = tmpPoint.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
