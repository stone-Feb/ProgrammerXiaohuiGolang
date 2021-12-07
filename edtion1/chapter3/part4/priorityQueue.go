package part4

import "fmt"

//最大堆实现最大优化队列

type priorityQueue struct {
	arr  []int
	size int
}

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{
		arr: make([]int, 0),
	}
}

//EnQueue 入队
func (p *priorityQueue) EnQueue(key int) {
	p.arr = append(p.arr, key)
	p.size++
	//最大堆实现
	p.upAdjust()
}

//DeQueue 出队
func (p *priorityQueue) DeQueue() int {
	if p.size <= 0 {
		panic("the queue is empty")
	}
	//获取堆顶元素
	head := p.arr[0]
	//最后一个元素移到堆顶
	p.arr[0] = p.arr[p.size-1]
	p.size--
	//下沉调整
	p.downAdjust()
	return head
}

func (p *priorityQueue) OutPut() {
	fmt.Println(p.arr)
}

//最大堆插入节点，上浮调整
func (p *priorityQueue) upAdjust() {
	childIndex := p.size - 1
	parentIndex := (childIndex - 1) / 2
	// temp保存插入的叶子节点值，用于最后的赋值
	temp := p.arr[childIndex]
	for childIndex > 0 && temp > p.arr[parentIndex] {
		p.arr[childIndex] = p.arr[parentIndex]
		childIndex, parentIndex = parentIndex, (parentIndex-1)/2
	}
	p.arr[childIndex] = temp
}

//最小堆删除节点，下沉操作
func (p *priorityQueue) downAdjust() {
	parentIndex := 0
	temp := p.arr[parentIndex]
	childIndex := 1
	for childIndex < p.size {
		// 如果有右孩子，且右孩子大于左孩子的值，则定位到右孩子
		if childIndex+1 < p.size && p.arr[childIndex+1] > p.arr[childIndex] {
			childIndex++
		}
		//如果父节点大于任何一个孩子的值，直接跳出
		if temp >= p.arr[childIndex] {
			break
		}
		//无需真正交换，单向赋值即可。堆顶永远是最大元素
		p.arr[parentIndex] = p.arr[childIndex]
		parentIndex, childIndex = childIndex, 2*childIndex+1
	}
	p.arr[parentIndex] = temp
}
