package stack

import (
	"fmt"
	"testing"
)

func TestInitStack(t *testing.T) {
	//初始化栈
	fmt.Println("初始化栈")
	stack := InitStack()
	stack.Push(1)
	stack.Push(3)
	stack.Push(4)
	stack.StackPrint()
	//pop操作
	fmt.Println("pop操作")
	fmt.Println(stack.Pop())
	stack.StackPrint()
	//查询top
	fmt.Println("查询top")
	fmt.Println(stack.Top())
	//检查是否为空
	fmt.Println("检查是否为空")
	fmt.Println(stack.IsEmpty())
	//获取栈大小
	fmt.Println("获取栈大小")
	stack.StackPrint()
	fmt.Println(stack.Size())
}
