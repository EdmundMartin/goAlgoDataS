package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testArray := []int{8, 12, 7, 1, 727, 127, 66}
	MergeSort(testArray, 0, len(testArray)-1)
	if !sameSliceInt(testArray, []int{1, 7, 8, 12, 66, 127, 727}) {
		t.Error("array incorrectly sorted")
	}
}