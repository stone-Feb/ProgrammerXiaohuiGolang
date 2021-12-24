package chapter5

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter5/part2"
	"ProgrammerXiaohuiGolang/edtion1/chapter5/part3"
	"fmt"
	"testing"
)

func TestLinkedListCycle( t *testing.T)  {
	node1 := part2.NewNode(1)
	node2 := part2.NewNode(2)
	node3 := part2.NewNode(3)
	node4 := part2.NewNode(4)
	node5 := part2.NewNode(5)
	node6 := part2.NewNode(6)
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node1
	fmt.Println(part2.IsCycle(node1))
}

func TestMinStack(t *testing.T)  {
	stack := part3.NewMinStack()
	stack.Push(4)
	stack.Push(9)
	stack.Push(7)
	stack.Push(3)
	stack.Push(8)
	stack.Push(5)
	fmt.Println(stack.GetMin())
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.GetMin())
}
