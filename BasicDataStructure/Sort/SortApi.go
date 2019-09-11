package Sort

type BaseSortApi interface {
	Comparable
	sort(array []int, value ...int)
	Sort(array []int)
}

type Comparable interface {
	less(a1, a2 int) bool
	swap(array []int, i, j int)
}
