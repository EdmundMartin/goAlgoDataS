package linkedlists

import "testing"

func isSameIntSlice(s1, s2 []int) bool {
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

func TestSinglyLinkedList_AddBack(t *testing.T) {
	newList := NewSinglyLinkedList()
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range values {
		newList.AddBack(val)
	}
	iterated := newList.IterateValues()
	if !isSameIntSlice(values, iterated) {
		t.Errorf("expected to find the same slice, %v, %v", values, iterated)
	}
}

func TestSinglyLinkedList_NodeAtIdx(t *testing.T) {
	newList := NewSinglyLinkedList()
	values := []int{1, 2, 3, 5, 6, 7}
	for _, val := range values {
		newList.AddBack(val)
	}
	target := newList.NodeAtIdx(3)
	if target.Value != 5 {
		t.Errorf("expected value %d, got %d", 5, target.Value)
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	newList := NewSinglyLinkedList()
	values := []int{0, 1, 2, 3, 4, 5, 6}
	for _, val := range values {
		newList.AddBack(val)
	}
	toRemove := newList.NodeAtIdx(2)
	newList.Remove(toRemove)
	iterated := newList.IterateValues()
	if !isSameIntSlice([]int{0, 1, 3, 4, 5, 6}, iterated) {
		t.Errorf("expected node at index two have been removed")
	}
}