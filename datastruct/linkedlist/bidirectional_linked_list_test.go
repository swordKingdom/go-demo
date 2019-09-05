package linkedlist

import (
	"fmt"
	"testing"
)

func TestBLinkedListHead(t *testing.T) {
	arr := make([]interface{}, 0)
	arr = append(arr, 1)
	arr = append(arr, 4)

	head := InitBLinkedList(arr)
	//插入链表数据
	head.InsertNode(2, 10)
	head.RangeReadOnly(func(node *Node) {
		fmt.Println(node.Val)
	})
	fmt.Println(head.GetNode(2))
	//删除链表数据
	head.DeleteNode(1)
	head.RangeReadOnly(func(node *Node) {
		fmt.Println(node.Val)
	})
	fmt.Println(head.GetNode(1))
}
