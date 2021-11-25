package part2

import "fmt"

//2.2 链表：
//特征：链式存储结构，顺序存取，
//分类：单向、静态、双向、有环

// MyLinkedLink 单链表实现
type MyLinkedLink struct {
	//头节点指针
	head *node
	//尾节点指针
	last *node
	//链表实际长度
	size int
}

// NewMyLinkedLink 构造函数
func NewMyLinkedLink() *MyLinkedLink {
	return &MyLinkedLink{}
}

// Insert 链表插入元素
// index 插入位置
// data 插入元素
func (l *MyLinkedLink) Insert(index, data int) {
	if index < 0 || index > l.size {
		panic("超出链表节点范围！")
	}
	//新建节点
	insertNode := &node{data: data}
	if l.size == 0 {
		// 空链表插入，头尾指针指向头节点
		l.head, l.last = insertNode, insertNode
	} else if index == 0 {
		//头部插入：1，新节点的next指针指向原来的头节点。2，head指针指向链表的头节点
		insertNode.next, l.head = l.head, insertNode
	} else if index == l.size {
		//尾插入：把最后一个节点的next指针指向新插入的节点即可
		l.last.next, l.last = insertNode, insertNode
	} else {
		//中间插入:1.新节点的next指针，指向插入位置的节点；2.插入位置的"直接前驱"的next指针，指向新节点
		prevNode := l.Get(index - 1)
		insertNode.next, prevNode.next = prevNode.next, insertNode
	}
	//增加链表长度
	l.size++
}

// Remove 链表删除元素
// index 删除位置
func (l *MyLinkedLink) Remove(index int) *node {
	if index < 0 || index >= l.size {
		panic("超出链表节点范围！")
	}
	var removeNode *node
	if index == 0 {
		//删除头节点:把链表的头节点设为与原来头节点的next指针即可
		removeNode, l.head = l.head, l.head.next
		//删除完了，尾节点置为空
		if l.size == 1 {
			l.last = &node{}
		}
	} else if index == l.size-1 {
		//删除尾节点：倒数第二个节点的next指针指为空
		prevNode := l.Get(index - 1)
		removeNode, prevNode.next, l.last = prevNode.next, &node{}, prevNode
	} else {
		//中间删除:把removeNode的"直接前驱"的next指针，指向removeNode的"直接后驱"
		prevNode := l.Get(index - 1)
		nextNode := prevNode.next.next
		prevNode.next = nextNode
		removeNode = prevNode.next
	}
	l.size--
	return removeNode
}

// Get 链表查询元素
// index 查找的位置
func (l *MyLinkedLink) Get(index int) *node {
	if index < 0 || index >= l.size {
		panic("超出链表节点范围！")
	}
	temp := l.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp
}

// Output 输出链表
func (l *MyLinkedLink) Output() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

//链表节点
type node struct {
	data int
	next *node
}
