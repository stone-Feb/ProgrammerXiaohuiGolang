package chapter3

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part2"
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part3"
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part4"
	"container/heap"
	"container/list"
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(4)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(5)
	l.PushBack(nil)
	l.PushBack(6)
	treeNode := part2.CreateBinaryTree(l)
	//前序遍历
	//part2.PreOrderTraversal(treeNode)
	//中序遍历
	//part2.InOrderTraversal(treeNode)
	//后序遍历
	part2.PostOrderTraversal(treeNode)
}

func TestTestNewBinaryTreeWithStack(t *testing.T) {
	//生成线性序列
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(4)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(5)
	l.PushBack(nil)
	l.PushBack(6)
	//生成树
	treeNode := part2.CreateBinaryTree(l)
	//栈实现遍历
	part2.PreOrderTraversalWithStack(treeNode)
}

func TestBinaryTreeTraversalLevel(t *testing.T) {
	//生成线性序列
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(4)
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(5)
	l.PushBack(nil)
	l.PushBack(6)
	//生成树
	treeNode := part2.CreateBinaryTree(l)
	part2.LevelOrderTraversal(treeNode)
}

func TestMinHeap(t *testing.T) {
	//最小堆,插入节点
	arr := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	part3.UpAdjust(arr)
	fmt.Println(arr)

	//无序二叉堆构建最小堆
	arr2 := []int{7, 1, 3, 10, 5, 2, 8, 9, 6}
	part3.BuildHeap(arr2)
	fmt.Println(arr2)
}

func TestPriorityQueue(t *testing.T) {
	queue := part4.NewPriorityQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.OutPut()
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	//fmt.Println(queue.DeQueue())
	//fmt.Println(queue.DeQueue())
	queue.OutPut()
}

func TestHeap(t *testing.T) {
	h := &part4.IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

//为基本类型添加方法
type inta []int

func (i *inta) push(j int) {
	*i = append(*i, j)
	fmt.Println(i)
}

func TestInt(t *testing.T)  {
	i := &inta{2,1,1}
	i.push(3)
}
