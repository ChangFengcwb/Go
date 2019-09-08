package Sort

type InMerge struct {
	aux []int
}

func (m *InMerge) less(a1, a2 int) bool {
	return a1 < a2
}

func (m *InMerge) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

//value [0]==lo value [1]==mid value [2]==hi
func (m *InMerge) sort(array []int, value ...int) {
	left, right := value[0], value[1]+1
	copy(m.aux[value[0]:value[2]+1], array[value[0]:value[2]+1])
	for k := value[0]; k <= value[2]; k++ {
		if left > value[1] {
			array[k] = m.aux[right]
			right++
		} else if right > value[2] {
			array[k] = m.aux[left]
			left++
		} else if m.less(m.aux[right], m.aux[left]) {
			array[k] = m.aux[right]
			right++
		} else {
			array[k] = m.aux[left]
			left++
		}
	}
}

func (m *InMerge) merge(array []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + ((hi - lo) >> 1)
	m.merge(array, lo, mid)
	m.merge(array, mid+1, hi)
	m.sort(array, lo, mid, hi)
}
func (m *InMerge) Sort(array []int) {
	if len(array) <= 0 {
		return
	}
	m.aux = make([]int, len(array))
	m.merge(array, 0, len(array)-1)
}
