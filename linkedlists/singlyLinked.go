package linkedlists

type SinglyLinkedNode struct {
	Value int
	Next *SinglyLinkedNode
}

func NewSinglyLinkedNode(value int)  *SinglyLinkedNode {
	return &SinglyLinkedNode{Value: value}
}

// AddFront adds a value to front of linked list and returns the head node of singly linked list
func AddFront(head *SinglyLinkedNode, value int) *SinglyLinkedNode {
	if head == nil {
		return NewSinglyLinkedNode(value)
	}
	newNode := NewSinglyLinkedNode(value)
	newNode.Next = head
	return newNode
}

// AddBack adds a value to the back of a linked list and returns the head of the signly linked list
func AddBack(head *SinglyLinkedNode, value int)  {
	// Edge case
	if head == nil {
		return
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewSinglyLinkedNode(value)
}

// RemoveNode and tell caller whether the node has been successfully removed
func RemoveNode(head *SinglyLinkedNode, node *SinglyLinkedNode) bool {
	if head == nil {
		return false
	}
	if head == node {
		head = head.Next
		return true
	}
	var prev *SinglyLinkedNode
	current := head
	for current != nil {
		if current == node && prev != nil {
			prev.Next = current.Next
			return true
		}
		prev = current
		current = current.Next
	}
	return false
}