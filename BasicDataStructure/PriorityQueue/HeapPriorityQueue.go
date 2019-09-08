package PriorityQueue

type HeapPriorityQueue struct {
	slice []int
	size  int
}

func NewHeapPriorityQueue() *HeapPriorityQueue {
	return &HeapPriorityQueue{
		slice: make([]int, 1),
		size:  0,
	}
}

func (this *HeapPriorityQueue) swim(k int) {
	for k > 0 && (k>>1) > 0 && this.less(k>>1, k) {
		this.swap(k>>1, k)
		k = k >> 1
	}
}

func (this *HeapPriorityQueue) sink(k int) {
	for 2*k < this.size {
		j := k << 1
		//先比较两个子树哪个小 ，跟大的换
		if j < this.size && this.less(j, j+1) {
			j++
		}
		if !this.less(k, j) {
			break
		}
		this.swap(k, j)
		k = j
	}
}
func (this *HeapPriorityQueue) Pop() int {
	var temp int
	if this.size > 0 {
		temp = this.slice[1]
		this.swap(1, this.size)
		this.slice = this.slice[:this.size]
		this.size--
		this.sink(1)
	} else {
		return 0
	}
	return temp

}

func (this *HeapPriorityQueue) Peek() int {
	if this.size > 0 {
		return this.slice[1]
	}
	return 0
}

func (this *HeapPriorityQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *HeapPriorityQueue) Insert(key int) {
	this.slice = append(this.slice, key)
	this.size++
	this.swim(this.size)
}

func (this *HeapPriorityQueue) Size() int {
	return this.size
}

func (this *HeapPriorityQueue) less(a1, a2 int) bool {
	return this.slice[a1] < this.slice[a2]
}

func (this *HeapPriorityQueue) swap(i, j int) {
	this.slice[i], this.slice[j] = this.slice[j], this.slice[i]
}
