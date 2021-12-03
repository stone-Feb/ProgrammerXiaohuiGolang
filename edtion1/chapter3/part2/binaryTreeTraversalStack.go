package part2

import "fmt"

//递归和栈都有回溯特性

//实现栈
type stack struct {
	element []*treeNode
}

func NewStack() *stack {
	return &stack{
		element: make([]*treeNode, 0),
	}
}

//入栈
func (s *stack) push(x *treeNode) {
	s.element = append(s.element, x)
}

//出栈
func (s *stack) pop() *treeNode {
	if len(s.element) == 0 {
		return nil
	}
	x := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return x
}

//判断空栈
func (s *stack) empty() bool {
	if len(s.element) == 0 {
		return true
	}
	return false
}

// PreOrderTraversalWithStack 栈实现二叉树前序遍历:根-左-右
func PreOrderTraversalWithStack(node *treeNode) {
	stack := NewStack()
	for node != nil || !stack.empty() {
		//迭代访问节点的左孩子，并入栈
		for node != nil {
			fmt.Println(node.data)
			stack.push(node)
			node = node.leftChild
		}
		//如果节点没有左孩子，则弹出栈顶节点，访问节点右孩子
		if !stack.empty() {
			node = stack.pop()
			node = node.rightChild
		}
	}
}
