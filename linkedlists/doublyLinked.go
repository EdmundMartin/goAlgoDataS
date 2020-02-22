package linkedlists

type DoubleNode struct {
	Value int
	Next *DoubleNode
	Prev *DoubleNode
}

func NewDoubleNode(val int) *DoubleNode {
	return &DoubleNode{Value: val}
}

type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	Size int
}

func NewDoubleLinkedList()  *DoubleLinkedList {
	return &DoubleLinkedList{Size: 0}
}

func (d *DoubleLinkedList) InsertFront(val int) {
	if d.Head == nil {
		node := &DoubleNode{Value: val}
		d.Head = node
		d.Tail = node
		d.Size++
	} else {
		node := &DoubleNode{Value: val}
		head := d.Head
		head.Prev = node
		node.Next = head
		d.Head = node
		d.Size++
	}
}

func (d *DoubleLinkedList) InsertEnd(val int) {
	if d.Tail == nil {
		node := &DoubleNode{Value: val}
		d.Head = node
		d.Tail = node
		d.Size++
	} else {
		node := &DoubleNode{Value: val}
		tail := d.Tail
		tail.Next = node
		node.Prev = tail
		d.Tail = node
		d.Size++
	}
}

func (d *DoubleLinkedList) Remove(val int) {
	current := d.Head
	for current != nil {
		if current.Value == val && current == d.Head {
			if d.Head == d.Tail {
				d.Head = nil
				d.Tail = nil
				return
			}
			d.Head = d.Head.Next
		} else if current.Value == val {
			if current == d.Tail {
				d.Tail = d.Tail.Prev
			}
			tempNext := current.Next
			tmpPrev := current.Prev
			tmpPrev.Next = tempNext
			tempNext.Prev = tmpPrev
		}
		current = current.Next
	}
}