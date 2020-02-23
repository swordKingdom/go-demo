package heap

import "sync"

//MaxHeap 最大堆
//TODO:实现同时支持基础类型：int,int64,float64等的最大堆
type MaxHeap struct {
	data      []int
	equalFunc func(a, b interface{})
	lock      *sync.RWMutex
}

//swap 交换数组元素
func (m *MaxHeap)swap(index1,index2 int) {
	m.data[index1] ,m.data[index2] =  m.data[index2] ,m.data[index1]
}

//adjustDown 将首部元素向下沉
func (m *MaxHeap) adjustDown(index int) {
	size :=len(m.data)
	parentIndex := index
	for parentIndex < size {
		leftChild := parentIndex <<1 + 1
		rightChild := leftChild +1
		maxIndex := leftChild
		if leftChild >size -1 {
			break
		}else if rightChild > size -1 {
			if m.data[leftChild] > m.data[parentIndex] {
				m.swap(leftChild,parentIndex)
			}
			break
		}else{
			if m.data[rightChild] > m.data[maxIndex] {
				maxIndex = rightChild
			}
			if m.data[maxIndex] < m.data[parentIndex] {
				break
			}
			m.swap(maxIndex,parentIndex)
			parentIndex = maxIndex
		}
	}
}

//adjustUp 向上上升
func (m *MaxHeap) adjustUp(index int) {
	childIndex := index
	if len(m.data) == 1 {
		return
	}
	for childIndex >0 {
		parentIndex  := (childIndex-1) >>1
		if m.data[parentIndex] < m.data[childIndex] {
			m.swap(parentIndex,childIndex)
			childIndex = parentIndex
		}else{
			break
		}
	}
	return
}

//Add 在堆中增加元素
func (m *MaxHeap) Add(val int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data = append(m.data, val)
	m.adjustUp(len(m.data) - 1)
}

//PollMin 取出堆中增加最大元素
func (m *MaxHeap) PollMax() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	size := len(m.data)
	if size == 0 {
		return 0
	}
	res := m.data[0]
	m.data[0] = m.data[size-1]
	m.data = m.data[:size-1]
	if len(m.data) != 0 {
		m.adjustDown(0)
	}
	return res
}

//InitMinHeap 初始化最小堆
func InitMaxHeap(arr []int) *MaxHeap {
	res := &MaxHeap{}
	res.data = make([]int, 0, len(arr))
	for _, e := range arr {
		res.Add(e)
	}
	return res
}
