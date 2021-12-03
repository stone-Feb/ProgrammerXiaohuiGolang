package part2

import "fmt"

//二叉树的层序遍历

//实现队列
type queue struct {
	element []*treeNode
}

func newQueue() *queue {
	return &queue{
		element: make([]*treeNode , 0),
	}
}

func (q *queue)inQueue(x *treeNode)  {
	q.element = append(q.element , x)
}

func (q *queue)enQueue() *treeNode {
	if len(q.element) == 0 {
		return nil
	}
	x := q.element[0]
	q.element = q.element[1:]
	return x
}

func (q *queue)empty() bool {
	if len(q.element) == 0 {
		return true
	}
	return false
}

//LevelOrderTraversal 层序遍历
func LevelOrderTraversal(node *treeNode) {
	queue := newQueue()
	//根入队
	queue.inQueue(node)
	//队列实现层序遍历
	for !queue.empty() {
		node = queue.enQueue()
		fmt.Println("root",node.data)
		//子节点入队
		if node.leftChild != nil {
			//fmt.Println("left",node.leftChild.data)
			queue.inQueue(node.leftChild)
		}
		if node.rightChild != nil {
			//fmt.Println("right",node.rightChild.data)
			queue.inQueue(node.rightChild)
		}
	}
}