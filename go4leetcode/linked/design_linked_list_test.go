package linked

import (
	"testing"
)

func TestConstructorLinked(t *testing.T) {
	obj := Constructor()
	//var param  int
	//obj.AddAtHead(1)
	//param = obj.Get(1)
	//fmt.Println(param)
	//
	//obj.AddAtTail(2)
	//param = obj.Get(2)
	//fmt.Println(param)
	//
	//obj.AddAtIndex(1, 3)
	//param = obj.Get(1)
	//fmt.Println(param)
	//
	//obj.DeleteAtIndex(2)
	//param = obj.Get(2)
	//fmt.Println(param)
	obj.AddAtHead(1)
	obj.AddAtTail(2)
	obj.AddAtIndex(1, 3)
	obj.DeleteAtIndex(2)
	obj.Get(2)
}
