package sorting

import (
	"testing"
)

func sameSliceInt(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx, val := range s1 {
		if s2[idx] != val {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	testArray := []int{12, 27, 26, 1, 8, 9, 10}
	result := QuickSort(testArray)
	if !sameSliceInt(result, []int{1, 8, 9, 10, 12, 26, 27}) {
		t.Error("array incorrectly sorted")
	}
}