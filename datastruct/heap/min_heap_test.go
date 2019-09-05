package heap

import (
	"fmt"
	"testing"
)

func TestInitMinHeap(t *testing.T) {
	arr := []int{2, 10, 3, 4, 5}
	//初始化最小堆
	heap := InitMinHeap(arr)
	fmt.Println(heap)
	//在最小堆添加元素
	heap.Add(1)
	fmt.Println(heap)
	//获取最小堆的元素
	data := heap.PollMin()
	fmt.Println(data)
	fmt.Println(heap)

}
