package problems

import "testing"

func TestTwoSum(t *testing.T) {
	test := []int{-2, 1, 2, 3, 3, 3, 4, 5, 6, 8}
	result := TwoSum(test, 6)
	expected := [][]int{
		[]int{3, 3},
		[]int{2, 4},
		[]int{1, 5},
		[]int{-2, 8},
	}
	if len(result) != len(expected) {
		t.Errorf("expected different number of unique results")
	}
	for idx, val := range expected {
		res := result[idx]
		if res[0] != val[0] && res[1] != val[1] {
			t.Errorf("expected a different result to be returned")
		}
	}
}
