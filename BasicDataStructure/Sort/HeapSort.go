package Sort

type HeapSort struct {
}

func (this *HeapSort) less(a1, a2 int) bool {
	return a1 < a2
}

func (this *HeapSort) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (this *HeapSort) sink(array []int, k, end int) {
	for 2*k < end {
		j := k << 1
		if j < end && this.less(array[j], array[j+1]) {
			j++
		}
		if !this.less(array[k], array[j]) {
			break
		}
		this.swap(array, k, j)
		k = j
	}
}

func (this *HeapSort) sort(array []int, value ...int) {
	n := len(array) - 1
	for i := n >> 1; i >= 0; i-- {
		this.sink(array, i, n)
	}
	for n > 1 {
		this.swap(array, 0, n)
		n--
		this.sink(array, 0, n)
	}
}

func (this *HeapSort) Sort(array []int) {
	if len(array) <= 1 {
		return
	}
	this.sort(array)
}
