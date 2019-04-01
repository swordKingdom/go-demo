package object_pool

import "fmt"

//对象池

type OpObject struct {
	Code int
}
func (v *OpObject)SayCode(i int){
	fmt.Errorf("my code is %v",v.Code+i)
}

type Pool chan *OpObject

func NewPool(total int) Pool {
	//此处一定要用带缓冲的channel，否则在同一个协程内使用不带缓冲的channel会出现死锁问题
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		op :=  new(OpObject)
		op.Code = i
		p <-  op
	}
	return p
}

func main(){
	p := NewPool(2)
	for i:=0 ;i<2 ;i++ {
		select {
		case obj,ok:= <-p:
			if ok {
				obj.SayCode(2)
			}
			p<- obj
		default:

			return
		}
	}
}