package Search

func BinarySearch(arr []int, key int) bool {
	if len(arr) <= 0 {
		return false
	}
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] < key {
			left = mid + 1
		} else if arr[mid] > key {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
