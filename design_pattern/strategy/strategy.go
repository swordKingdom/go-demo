package main

import "fmt"

//策略模式
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Subtraction struct{}

func (Subtraction) Apply(lval, rval int) int {
	return lval - rval
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}


func main(){
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5)) // 8
	sub := Operation{Subtraction{}}
	fmt.Println(sub.Operate(3, 5)) // -2
}