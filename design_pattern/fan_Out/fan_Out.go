package main

//扇出模式，用于消息分流
func Split(ch chan int, n int) []chan int {
	cs := make([]chan int,0)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}
	distributeToChannels := func(chTmp <-chan int, cs []chan int) {
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}
