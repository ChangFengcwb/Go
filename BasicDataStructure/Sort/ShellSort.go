package Sort

type Shell struct {
}

func (s *Shell) Sort(array []int) {
	if len(array) <= 0 {
		return
	}
	s.sort(array)
}

func (s *Shell) less(a1, a2 int) bool {
	return a1 < a2
}

func (s *Shell) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (s *Shell) sort(array []int, value ...int) {
	h := 1
	for h < h/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(array); i++ {
			for j := i; j >= h && s.less(array[j], array[j-h]); j -= h {
				s.swap(array, j, j-h)
			}
		}
		h = h / 3
	}
}
