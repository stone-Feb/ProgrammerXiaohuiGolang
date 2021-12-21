package part3

//栈实现
type stack struct {
	element []map[string]int
}

func newStack() *stack {
	return &stack{
		element: []map[string]int{},
	}
}

func (s *stack) pop() map[string]int {
	if s.empty() {
		panic("stack empty")
	}
	length := len(s.element)
	ele := s.element[length-1]
	s.element = s.element[:length-1]
	return ele
}

func (s *stack) push(ele map[string]int) {
	s.element = append(s.element, ele)
}

func (s *stack) empty() bool {
	if len(s.element) == 0 {
		return true
	}
	return false
}

// QuickSortWithStack 栈实现快速排序
func QuickSortWithStack(arr []int, startIndex, endIndex int) {
	//栈代替递归
	quickSortStack := newStack()
	//入栈
	rootParam := map[string]int{}
	rootParam["startIndex"] = startIndex
	rootParam["endIndex"] = endIndex
	quickSortStack.push(rootParam)

	//循环结束：栈为空
	for !quickSortStack.empty() {
		//栈顶出栈
		param := quickSortStack.pop()
		//获取基准元素
		pivotIndex := partitionV2(arr, param["startIndex"], param["endIndex"])
		//分治法：根据基准元素分成两部分，分别进行单边循环
		if param["startIndex"] < pivotIndex-1 {
			leftParam := map[string]int{}
			leftParam["startIndex"] = param["startIndex"]
			leftParam["endIndex"] = pivotIndex - 1
			quickSortStack.push(leftParam)
		}
		if pivotIndex+1 < param["endIndex"] {
			rightParam := map[string]int{}
			rightParam["startIndex"] = pivotIndex + 1
			rightParam["endIndex"] = param["endIndex"]
			quickSortStack.push(rightParam)
		}
	}
}
