package main

import "fmt"

//代理模式
type IData interface {
	Process(opt string) error
}

type DataImpl struct {
	opt string
}

func (v *DataImpl) Process(opt string) error {
	fmt.Printf("op, %s", v.opt)
	return nil
}

type ProxyData struct {
	data *DataImpl
}

func (v *ProxyData) Process(opt string) error {
	if v.data == nil {
		v.data = new(DataImpl)
	}
	if opt == "Run" {
		v.data.Process(opt)
	}
}
