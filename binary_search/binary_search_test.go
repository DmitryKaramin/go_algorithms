package binary_search

import (
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	myArr := []int{1, 3, 5, 7, 9}
	expResult := 1
	var expBool bool

	recResult, recBool := binarySearch(myArr, 3)
	if recResult != expResult && recBool != expBool {
		t.Fatalf("Received %v, %v not equals %v, %v ", recResult, recBool, expResult, expBool)
	}
	log.Printf("Found %v in %v", recResult, myArr)
}