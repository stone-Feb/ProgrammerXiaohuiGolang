package part3

import "fmt"

//2.3 栈和队列
//逻辑结构：线性结构。物理实现：数组（顺序存储）或链表（链式存储）
//stack 栈：First In Last Out 先入后出。最早入栈元素为栈底bottom，最迟入栈的元素为栈顶top
//queue 队列：First In First Out 先入先出。最早入队元素为队头front，最迟入队的元素为队尾rear

//循环队列
//1.数组实现的队列可以采用循环队列的方式来维持队列容量的恒定
//2.当（队尾下标+1）%数组长度 = 队头下标时，代表队列已满。需要队尾指针指向的位置永远空出1位，所以循环队列最大容量比数组长度小1

type myQueue struct {
	array []int
	//队头下标
	front int
	//队尾下标
	rear int
}

func NewMyQueue(capacity int) *myQueue {
	return &myQueue{
		array: make([]int, capacity),
	}
}

// EnQueue 入队
// element 入队元素
func (q *myQueue) EnQueue(element int) {
	if (q.rear+1)%len(q.array) == q.front {
		panic("队列已满")
	}
	q.array[q.rear] = element
	q.rear = (q.rear + 1) % len(q.array)
}

// DeQueue 出队
func (q *myQueue) DeQueue() int {
	if q.rear == q.front {
		panic("队列已空")
	}
	deQueueElement := q.array[q.front]
	q.front = (q.front + 1) % len(q.array)
	return deQueueElement
}

func (q *myQueue) Output() {
	for i := q.front ; i != q.rear ; i = (i+1)%len(q.array) {
		fmt.Println(q.array[i])
	}
}
