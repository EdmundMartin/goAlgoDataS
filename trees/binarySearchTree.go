package trees

import "sort"

type BST struct {
	Value int
	Left *BST
	Right *BST
}

func NewBST(val int) *BST {
	return &BST{Value: val}
}

func (bst *BST) Insert(val int) {
	currentNode := bst
	if val < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = NewBST(val)
		} else {
			currentNode.Left.Insert(val)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = NewBST(val)
		} else {
			currentNode.Right.Insert(val)
		}
	}
}

func (bst *BST) Contains(val int) bool {
	currentNode := bst
	for currentNode != nil {
		if val < currentNode.Value {
			currentNode = currentNode.Left
		} else if val == currentNode.Value {
			return true
		} else {
			currentNode = currentNode.Right
		}
	}
	return false
}

func (bst *BST) Remove(val int, parent *BST) {
	currentNode := bst
	for currentNode != nil {
		if val < currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Left
		} else if val > currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Right
		} else {
			// TODO Clean up deletion logic
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Value = currentNode.Right.getMinValue()
				currentNode.Right.Remove(currentNode.Value, currentNode)
			} else if parent == nil {
				if currentNode.Left != nil {
					currentNode.Value = currentNode.Left.Value
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Value = currentNode.Right.Value
					currentNode.Left = currentNode.Right.Left
					currentNode.Right = currentNode.Right.Right
				} else {
					currentNode.Value = -1
				}
			} else if parent.Left == currentNode {
				var child *BST
				if currentNode.Left != nil {
					child = currentNode.Left
				} else {
					child = currentNode.Right
				}
				parent.Left = child
			} else if parent.Right == currentNode {
				var child *BST
				if currentNode.Left != nil {
					child = currentNode.Left
				} else {
					child = currentNode.Right
				}
				parent.Right = child
			}
			break
		}
	}
}


func (bst *BST) getMinValue()  int {
	currentNode := bst
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Value
}


func BuildBalancedBST(values []int) *BST {
	if len(values) == 0 {
		return nil
	}
	if len(values) == 1 {
		return NewBST(values[0])
	}
	sort.Ints(values)
	middle := len(values) / 2
	root := NewBST(values[middle])
	root.Left = BuildBalancedBST(values[:middle])
	root.Right = BuildBalancedBST(values[middle+1:])
	return root
}


func InOrder(bst *BST, values *[]int) {
	if bst == nil {
		return
	}
	InOrder(bst.Left, values)
	*values = append(*values, bst.Value)
	InOrder(bst.Right, values)
}