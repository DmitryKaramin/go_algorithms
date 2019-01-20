package quick_sort


func quickSort(arr []int) []int {
	var less []int
	var greater []int
	var pivot int

	if len(arr) <= 1 {
		return arr
	} else {
		pivot = arr[len(arr) / 2]
		for _, v := range arr {
			if v < pivot {
				less = append(less, v)
			} else if v > pivot{
				greater = append(greater, v)
			}
		}

	}

	b := append(quickSort(less), []int{pivot}...)
	return append(b, quickSort(greater)...)
}