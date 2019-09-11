package DataStructure

type queue interface {
	//入队
	Push(interface{})
	//出队 如果队列为空 则返回nil
	Pop() interface{}
	//返回头元素 但队列不改变
	FrontElement() interface{}
	//判断是否为空 如果为空则返回true  否则返回false
	IsEmpty() bool
	//长度
	Len() int
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (q *Queue) Push(value interface{}) {
	if q.head == nil && q.tail == nil {
		q.head = newNode(value)
		q.tail = q.head
	} else {
		q.tail.next = newNode(value)
		q.tail = q.tail.next
	}
	q.len++
}

func (q *Queue) Pop() interface{} {
	if q.head == nil && q.tail == nil {
		return nil
	} else {
		temp := q.head.value
		q.head = q.head.next
		q.len--
		if q.len == 0 {
			q.tail = q.head
		}
		return temp
	}
}

func (q *Queue) FrontElement() interface{} {
	if q.head == nil && q.tail == nil {
		return nil
	}
	return q.head.value
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Len() int {
	return q.len
}
