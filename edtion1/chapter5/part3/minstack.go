package part3

// 最小栈实现
type intStack struct {
	element []int
}

func newStack() *intStack {
	return &intStack{
		element: []int{},
	}
}

func (s *intStack) push(ele int) {
	s.element = append(s.element, ele)
}

func (s *intStack) pop() int {
	length := len(s.element)
	if length == 0 {
		panic("stack is empty")
	}
	x := s.element[length-1]
	s.element = s.element[:length-1]
	return x
}

func (s *intStack) peek() int {
	length := len(s.element)
	return s.element[length-1]
}

func (s *intStack) empty() bool {
	if len(s.element) == 0 {
		return true
	}
	return false
}

type MinStack struct {
	mainStack intStack
	minStack  intStack
}

func NewMinStack() *MinStack {
	return &MinStack{
		mainStack: intStack{},
	}
}

func (m *MinStack) Push(element int) {
	m.mainStack.push(element)
	//辅助栈为空 || 主栈新元素 <= 辅助栈栈顶 ，进辅助栈
	if m.minStack.empty() || element <= m.minStack.peek() {
		m.minStack.push(element)
	}
}

func (m *MinStack) Pop() int {
	//主出栈元素 == 辅出栈元素 ， 辅栈顶出栈
	if m.mainStack.peek() == m.minStack.peek() {
		m.minStack.pop()
	}
	return m.mainStack.pop()
}

func (m *MinStack) GetMin() int{
	if m.mainStack.empty() {
		panic("stack is empty")
	}
	return m.minStack.peek()
}
