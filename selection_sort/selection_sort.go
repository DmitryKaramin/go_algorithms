package selection_sort

type MyArray struct {
	arr []int
}

func (arr *MyArray) Pop(index int){
	arr.arr[index], arr.arr = arr.arr[len(arr.arr)-1], arr.arr[:len(arr.arr)-1]
}

func findSmallest(arr *MyArray) int{
	smallest := arr.arr[0]
	smallestIndex := 0

	for i := 1; i <= len(arr.arr) - 1 ; i++ {
		if arr.arr[i] < smallest {
			smallest = arr.arr[1]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int{
	myArray := MyArray{arr:arr}

	var newArr []int
	for range myArray.arr {
		smallest := findSmallest(&myArray)
		newArr = append(newArr, myArray.arr[smallest])
		myArray.Pop(smallest)
	}
	return newArr
}