package search

import "testing"

func TestBinarySearch(t *testing.T) {
	targetArray := []int{1, 6, 8, 12, 16, 17, 19, 39, 45, 55, 64, 68, 79, 81}
	result := BinarySearch(targetArray, 67)
	if result != -1 {
		t.Error("did not expect to find value in array")
	}
	result = BinarySearch(targetArray, 19)
	if result != 6 {
		t.Error("incorrect idx for result")
	}
}