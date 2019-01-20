package binary_search

func binarySearch(l []int, i int) (int, bool) {
	low := 0
	high := len(l) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := l[mid]
		if guess == i {
			return mid, true
		}
		if guess > i {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
