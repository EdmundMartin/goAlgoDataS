package trees

import (
	"sort"
	"testing"
)


func sameIntSlice(s1, s2 []int) bool {
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

func TestInOrder(t *testing.T) {
	testArray := []int{1, 12, 151, 15, 24, 13, 246, 173, 87, 7}
	tree := BuildBalancedBST(testArray)
	result := []int{}
	InOrder(tree, &result)
	sort.Ints(testArray)
	if !sameIntSlice(testArray, result) {
		t.Error("expected inorder traversal to be sorted")
	}
}