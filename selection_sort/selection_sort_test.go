package selection_sort

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)


func TestSelectionSort(t *testing.T) {
	myArr := []int{5, 3, 6, 2, 10}
	expArr := []int{2, 3 ,5 ,6, 10}
	fmt.Print(reflect.TypeOf(expArr))

	recArr := selectionSort(myArr)

	if !reflect.DeepEqual(expArr, recArr) {
		log.Fatalf("Received %v not equals %v", recArr, expArr)
	}
	log.Printf("Slece %v equals %v", recArr, expArr)
}