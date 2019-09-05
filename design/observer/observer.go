package main

import (
	"fmt"
	"sync"
)

type (
	//事件结构体
	Event struct {
		Data int64
	}
	//观察者
	IObserver interface {
		OnNotify(Event)
	}
	//通知者
	INotifier interface {
		Register(name string, o IObserver)
		Deregister(name string)
		Notify(Event)
	}
)

type Notifier struct {
	observers sync.Map
}

func (v *Notifier) Register(name string, o IObserver) {
	if _, ok := v.observers.Load(name); !ok {
		v.observers.Store(name, o)
	}
}

func (v *Notifier) Deregister(name string) {
	v.observers.Delete(name)
}

func (v *Notifier) Notify(event Event) {
	v.observers.Range(func(key, value interface{}) bool {
		observer, ok := value.(IObserver)
		if !ok {
			return true
		}
		observer.OnNotify(event)
		return true
	})
}

type Observer struct {
	Name string
}

func (v *Observer) OnNotify(e Event) {
	fmt.Printf(" %v get event %v ", v.Name, e.Data)
	fmt.Println()
}

var notifier = new(Notifier)

func main() {
	observer1 := &Observer{Name: "一号"}
	observer2 := &Observer{Name: "二号"}
	notifier.Register(observer1.Name, observer1)
	notifier.Register(observer2.Name, observer2)
	notifier.Notify(Event{Data: 1})
}
