package part2

//如何判断链表有环
type node struct {
	data int
	Next *node
}

func NewNode(data int) *node {
	return &node{
		data: data,
	}
}

func IsCycle(head *node) bool {
	p1 := head
	p2 := head
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}

// 判断判断环长度；入环节点
// 1. 环长度：相遇后继续循环，第二次相遇的前进次数就是环长(追及问题，快的比慢的套环)= 速度差 * 前进次数
// 2. 入环节点问题：p1 = D+s1 ;p2 = D+s1+n(s1+s2) so D = (n-1)(s1+s2) + s2 answer : p1放头节点位置p2放相遇点，每次都前进一步，相遇节点就是入环节点
