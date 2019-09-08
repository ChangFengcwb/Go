package Sort

import (
	"math/rand"
)

type QuickSort struct {
}

func (q *QuickSort) less(a1, a2 int) bool {
	return a1 < a2
}

func (q *QuickSort) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

//value [0]==lo  value [1]==hi
func (q *QuickSort) sort(array []int, value ...int) {
	if value[1] <= value[0] {
		return
	}
	lt, i, gt, key := value[0], value[0]+1, value[1], rand.Intn(value[1]-value[0])+value[0]
	array[key], array[lt] = array[lt], array[key]
	key = array[lt]
	for i <= gt {
		if q.less(array[i], key) || array[i] == key {
			q.swap(array, i, lt)
			lt++
			i++
		} else {
			q.swap(array, i, gt)
			gt--
		}
	}
	q.sort(array, value[0], lt-1)
	q.sort(array, gt+1, value[1])
}

func (q *QuickSort) Sort(array []int) {
	if len(array) <= 0 {
		return
	}
	q.sort(array, 0, len(array)-1)
}
