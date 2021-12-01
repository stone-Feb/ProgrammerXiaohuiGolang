package part2

import (
	"container/list"
	"fmt"
)

//遍历方式（父节点和孩子节点的访问顺序）
//深度优先遍历：前序 中序 后序
//广度优先遍历：层序

//遍历二叉树
type treeNode struct {
	data       interface{}
	leftChild  *treeNode
	rightChild *treeNode
}

// CreateBinaryTree 构建二叉树
// l 输入序列
func CreateBinaryTree(l *list.List) *treeNode {
	var node *treeNode
	if l == nil || l.Len() == 0 {
		return nil
	}
	//删除并放回第一个list元素
	firstList := l.Front()
	data := firstList.Value
	l.Remove(firstList)
	//递归创建左右孩子节点
	//data == nil 说明该节点不存在
	if data != nil {
		node = &treeNode{
			data: data,
		}
		node.leftChild = CreateBinaryTree(l)
		node.rightChild = CreateBinaryTree(l)
	}
	return node
}

// PreOrderTraversal 二叉树前序遍历
// node 二叉树节点
func PreOrderTraversal(node *treeNode) {
	if node == nil {
		return
	}
	//前序：根-左-右
	fmt.Println(node.data)
	PreOrderTraversal(node.leftChild)
	PreOrderTraversal(node.rightChild)
}

// InOrderTraversal 二叉树中序遍历
// node 二叉树节点
func InOrderTraversal(node *treeNode)  {
	if node == nil {
		return
	}
	//中序：左-根-右
	InOrderTraversal(node.leftChild)
	fmt.Println(node.data)
	InOrderTraversal(node.rightChild)
}

// PostOrderTraversal 二叉树后序遍历
// node 二叉树节点
func PostOrderTraversal(node *treeNode)  {
	if node == nil {
		return
	}
	//后序：左-右-根
	PostOrderTraversal(node.leftChild)
	PostOrderTraversal(node.rightChild)
	fmt.Println(node.data)
}