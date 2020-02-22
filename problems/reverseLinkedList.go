package problems

import "github.com/EdmundMartin/goAlgoDataS/linkedlists"

func ReverseLinkedList(list *linkedlists.Node) *linkedlists.Node {
	var first *linkedlists.Node
	second := list
	for second != nil {
		third := second.Next
		second.Next = first
		first = second
		second = third
	}
	return first
}