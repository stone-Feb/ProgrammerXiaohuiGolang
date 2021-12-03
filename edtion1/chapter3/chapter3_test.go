package chapter3

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter3/part2"
	"container/list"
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