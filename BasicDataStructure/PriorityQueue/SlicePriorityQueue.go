package PriorityQueue

import "sort"

type slicePriorityQueue struct {
	slice []int
	size  int
	max   int
}

func NewSlicePriorityQueue() *slicePriorityQueue {
	return &slicePriorityQueue{
		slice: make([]int, 0),
		size:  0,
		max:   0,
	}
}

func (s *slicePriorityQueue) Pop() int {
	temp := 0
	if s.size > 1 {
		temp = s.slice[len(s.slice)-1]
		s.slice = s.slice[:len(s.slice)-1]
		s.max = s.slice[len(s.slice)-1]
		s.size--
	} else {
		return 0
	}
	return temp
}

func (s *slicePriorityQueue) Peek() int {
	return s.max
}

func (s *slicePriorityQueue) IsEmpty() bool {
	return s.size == 0
}

func (s *slicePriorityQueue) Insert(key int) {
	s.slice = append(s.slice, key)
	s.size++
	sort.Ints(s.slice)
	s.max = s.slice[len(s.slice)-1]
}

func (s *slicePriorityQueue) Size() int {
	return s.size
}

func (s *slicePriorityQueue) less(a1, a2 int) bool {
	return a1 < a2
}

func (s *slicePriorityQueue) swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}
