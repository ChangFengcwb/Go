package Sort

type Selection struct {
}

func (s *Selection) Sort(array []int) {
	if len(array) <= 0 {
		return
	}
	s.sort(array)
}

func (s *Selection) less(a1, a2 int) bool {
	return a1 < a2
}

func (s *Selection) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (s *Selection) sort(array []int, value ...int) {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i; j < len(array); j++ {
			if s.less(array[j], array[min]) {
				min = j
			}
		}
		s.swap(array, min, i)
	}
}
