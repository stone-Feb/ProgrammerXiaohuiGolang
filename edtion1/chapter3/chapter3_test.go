package chapter3

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part2"
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part3"
	"container/list"
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T)  {
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

func TestTestNewBinaryTreeWithStack(t *testing.T)  {
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

func TestBinaryTreeTraversalLevel(t *testing.T)  {
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

func TestMinHeap(t *testing.T)  {
	arr := []int{1,3,2,6,5,7,8,9,10,0}
	part3.UpAdjust(arr)
	fmt.Println(arr)

	arr2 := []int{7,1,3,10,5,2,8,9,6}
	part3.BuildHeap(arr2)
	fmt.Println(arr2)
}