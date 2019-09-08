package Sort

type BaseSortApi interface {
	less(a1, a2 int) bool
	swap(array []int, i, j int)
	sort(array []int, value ...int)
	Sort(array []int)
}
