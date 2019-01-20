package quick_sort

import (
	"log"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	myArr := []int{10, 5, 2, 3}
	expArr := []int{2, 3, 5, 10}
	recArr := quickSort(myArr)

	if !reflect.DeepEqual(expArr, recArr) {
		log.Fatalf("Received %v not as expected %v", recArr, expArr)
	}
	log.Printf("%v equals %v", expArr, recArr)

}
