package main

import (
	"sync"
	"fmt"
)
//单例模式
type singleton map[string]string

var (
	//sync.Once函数是在所有程序中，不过调用多少次，只会执行一次，具体明细查看sync.Once的用法
	once sync.Once
	instance singleton
)

func NewSingleton() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func main(){
	s := NewSingleton()
	s["this"] = "that"
	s2 :=NewSingleton()
	fmt.Println("This is ", s2["this"])
}