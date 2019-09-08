package Sort

type Insert struct {
}

func (s *Insert) Sort(array []int) {
	if len(array) <= 0 {
		return
	}
	s.sort(array)
}

func (s *Insert) less(a1, a2 int) bool {
	return a1 < a2
}

func (s *Insert) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (s *Insert) sort(array []int, value ...int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && s.less(array[j], array[j-1]); j-- {
			s.swap(array, j, j-1)
		}
	}
}
