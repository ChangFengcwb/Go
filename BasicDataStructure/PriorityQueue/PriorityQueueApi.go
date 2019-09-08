package PriorityQueue

type PriorityQueueApi interface {
	Pop() int
	Peek() int
	IsEmpty() bool
	Insert(key int)
	Size() int
	less(a1, a2 int) bool
	swap(i, j int)
}
