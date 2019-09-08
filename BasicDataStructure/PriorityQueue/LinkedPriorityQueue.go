package PriorityQueue

import "container/list"

type LinkedPriorityQueue struct {
	list *list.List
	max  int
}

func NewLinkedPriorityQueue() *LinkedPriorityQueue {
	return &LinkedPriorityQueue{
		list: list.New(),
		max:  0,
	}
}

func (l *LinkedPriorityQueue) Pop() int {
	if !l.IsEmpty() {
		temp := l.list.Front()
		l.list.Remove(temp)
		a := l.list.Front()
		for a != nil {
			if a.Value.(int) > l.max {
				l.max = a.Value.(int)
			}
			a = a.Next()
		}
		return temp.Value.(int)
	}
	return 0
}

func (l *LinkedPriorityQueue) Peek() int {
	if !l.IsEmpty() {
		return l.max
	}
	return 0
}

func (l *LinkedPriorityQueue) IsEmpty() bool {
	return l.list.Len() == 0
}

func (l *LinkedPriorityQueue) Insert(key int) {
	if key > l.max {
		l.max = key
		l.list.PushFront(key)
	} else {
		l.list.PushBack(key)
	}
}

func (l *LinkedPriorityQueue) Size() int {
	return l.list.Len()
}

func (l *LinkedPriorityQueue) less(a1, a2 int) bool {
	return a1 < a2
}

func (l *LinkedPriorityQueue) swap(i, j int) {
}
