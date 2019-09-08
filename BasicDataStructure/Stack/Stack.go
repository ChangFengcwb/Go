package Stack

type stack interface {
	//入队
	Push(interface{})
	//出队 如果队列为空 则返回nil
	Pop() interface{}
	//返回头元素 但队列不改变
	IsEmpty() bool
	//长度
	Len() int
}

type node struct {
	value interface{}
	pre   *node
	next  *node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		pre:   nil,
		next:  nil,
	}
}

type Stack struct {
	down *node
	top  *node
	len  int
}

func NewStack() *Stack {
	return &Stack{
		down: nil,
		top:  nil,
		len:  0,
	}
}

func (s *Stack) Push(data interface{}) {
	if s.top == nil && s.down == nil {
		s.top = newNode(data)
		s.down = s.top
	} else {
		s.down.next = newNode(data)
		s.down.next.pre = s.down
		s.down = s.down.next
	}
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.down == nil {
		return nil
	}
	temp := s.down.value
	if s.len == 1 {
		s.top = nil
		s.down = nil
		s.len--
	} else {
		s.down = s.down.pre
		s.len--
	}
	return temp
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Len() int {
	return s.len
}
