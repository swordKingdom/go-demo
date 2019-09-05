package queue

import (
	"fmt"
	"testing"
)

func TestInitQueue(t *testing.T) {
	queue := InitQueue(10)
	queue.InQueue(1)
	queue.InQueue(11111)
	queue.Traverse(func(a interface{}) {
		fmt.Println(a)
	})
}
