package linkedlists

type Node struct {
	Value int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{Value:value}
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{nil, nil}
}

func (s *SinglyLinkedList) AddFront(value int) {
	if s.Head == nil {
		node := NewNode(value)
		s.Head = node
		s.Tail = node
	} else {
		head := NewNode(value)
		head.Next = s.Head
		s.Head = head
	}
}

func (s *SinglyLinkedList) AddBack(value int) {
	if s.Tail == nil {
		node := NewNode(value)
		s.Head = node
		s.Tail = node
	} else {
		tail := s.Tail
		newNode := NewNode(value)
		tail.Next = newNode
		s.Tail = newNode
	}
}

func (s *SinglyLinkedList) Remove(node *Node) {
	if node == s.Head {
		s.Head = s.Head.Next
		return
	}
	var prev *Node
	current := s.Head
	for current != nil {
		if current == node && prev != nil {
			prev.Next = current.Next
			if current == s.Tail {
				s.Tail = prev
			}
		}
		prev = current
		current = current.Next
	}
}

func (s *SinglyLinkedList) IterateValues() []int {
	values := []int{}
	current := s.Head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func (s *SinglyLinkedList) NodeAtIdx(idx int) *Node {
	currentIdx := 0
	current := s.Head
	for current != nil && currentIdx <= idx {
		if currentIdx == idx {
			return current
		}
		current = current.Next
		currentIdx++
	}
	return nil
}


func SliceToSLL(slice []int) *SinglyLinkedList {
	ll := NewSinglyLinkedList()
	for _, val := range slice {
		ll.AddBack(val)
	}
	return ll
}