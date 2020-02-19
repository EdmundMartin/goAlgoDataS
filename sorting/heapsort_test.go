package sorting

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testArray := []int{12, 27, 26, 1, 8, 9, 10}
	result := HeapSort(testArray)
	fmt.Println(result)
	if !sameSliceInt(result, []int{1, 8, 9, 10, 12, 26, 27}) {
		t.Error("array incorrectly sorted")
	}
}