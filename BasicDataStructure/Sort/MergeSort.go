package Sort

type Merge struct {
}

func (m *Merge) less(a1, a2 int) bool {
	return a1 < a2
}

func (m *Merge) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

//value [0]==lo value [1]==mid value [2]==hi
func (m *Merge) sort(array []int, value ...int) {
	left, right := value[0], value[1]+1
	temp := make([]int, value[2]-value[0]+1)
	i := 0
	for left <= value[1] && right <= value[2] {
		if m.less(array[left], array[right]) || array[left] == array[right] {
			temp[i] = array[left]
			i++
			left++
		} else {
			temp[i] = array[right]
			i++
			right++
		}
	}
	for left <= value[1] {
		temp[i] = array[left]
		i++
		left++
	}
	//copy会增加大约0.01秒的运行时长
	//if left <= value[1] {
	//	copy(temp[i:], array[left:])
	//}
	//if right <= value[2] {
	//	copy(temp[i:], array[right:])
	//}
	for right <= value[2] {
		temp[i] = array[right]
		i++
		right++
	}
	copy(array[value[0]:], temp)
}

func (m *Merge) merge(array []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + ((hi - lo) >> 1)
	m.merge(array, lo, mid)
	m.merge(array, mid+1, hi)
	m.sort(array, lo, mid, hi)
}
func (m *Merge) Sort(array []int) {

	if len(array) <= 0 {
		return
	}
	m.merge(array, 0, len(array)-1)
}
