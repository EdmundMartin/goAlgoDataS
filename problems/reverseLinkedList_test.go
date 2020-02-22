package problems

import (
	"github.com/EdmundMartin/goAlgoDataS/linkedlists"
	"testing"
)


func createLinkedList(array []int) *linkedlists.Node {
	head := &linkedlists.Node{Value: array[0]}
	for i := 1; i < len(array); i++ {
		next := &linkedlists.Node{Value: array[i]}
		next.Next = head
		head = next
	}
	return head
}

func iterateLinkedList(head *linkedlists.Node) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

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

func TestReverseLinkedList(t *testing.T) {
	test := []int{5, 4, 3, 2, 1}
	ll := createLinkedList(test)
	result := ReverseLinkedList(ll)
	if !sameIntSlice(iterateLinkedList(result), test) {
		t.Error("expected input to match output")
	}
}