package heap

//MinHeap 最小堆使用
//todo:实现同时支持基础类型：int,int64,float64等的最小堆
type MinHeap struct {
	data      []int
	equreFunc func(a, b interface{})
}

//adjustDown 向下沉淀
func (m *MinHeap) adjustDown(index int) {
	lenth := len(m.data)
	left := index*2 + 1
	right := left + 1
	min := left
	if right > lenth-1 {
		return
	} else {
		if m.data[right] < m.data[min] {
			min = right
		}
	}
	if m.data[min] > m.data[index] {
		return
	}
	m.data[min], m.data[index] = m.data[index], m.data[min]
	m.adjustDown(min)
}

//adjustUp 向上上升
func (m *MinHeap) adjustUp(index int) {
	parent := (index - 1) / 2
	if parent < 0 {
		return
	}
	if m.data[parent] > m.data[index] {
		m.data[parent], m.data[index] = m.data[index], m.data[parent]
		m.adjustUp(parent)
	}
	return
}

//Add 在堆中增加元素
func (m *MinHeap) Add(val int) {
	m.data = append(m.data, val)
	m.adjustUp(len(m.data) - 1)
}

//PollMin 取出堆中增加最小元素
func (m *MinHeap) PollMin() int {
	dataLen := len(m.data)
	if dataLen == 0 {
		return 0
	}
	res := m.data[0]
	m.data[0] = m.data[dataLen-1]
	m.data = m.data[:dataLen-1]
	if len(m.data) != 0 {
		m.adjustDown(0)
	}
	return res
}

//InitMinHeap 初始化最小堆
func InitMinHeap(arr []int) *MinHeap {
	res := &MinHeap{}
	res.data = make([]int, 0, len(arr))
	for _, e := range arr {
		res.Add(e)
	}
	return res
}
