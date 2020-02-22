package caches

import "fmt"

type MissingKey struct {
	Key string
}

func (m MissingKey) Error() string {
	return fmt.Sprintf("no key: %s found", m.Key)
}

type node struct {
	Key   string
	Value string
	Next  *node
	Prev  *node
}

func newNode(key string, value string) *node {
	return &node{Key: key, Value: value}
}

func (n *node) String() string {
	return fmt.Sprintf("<Node: Key: %s, Value: %s", n.Key, n.Value)
}

type linkedList struct {
	Head *node
	Tail *node
}

func newLinkedList() *linkedList {
	return &linkedList{nil, nil}
}

func (l *linkedList) AddFront(n *node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		head := n
		head.Next = l.Head
		l.Head.Prev = head
		l.Head = head
	}
}

func (l *linkedList) Remove(n *node) {
	if n == l.Head {
		if n.Next != nil {
			l.Head = n.Next
		} else {
			l.Head = nil
		}
		return
	}
	if n == l.Tail {
		if n.Prev != nil {
			l.Tail = n.Prev
		}
	}
	nextNode := n.Next
	prevNode := n.Prev
	if nextNode != nil {
		nextNode.Prev = prevNode
	}
	if prevNode != nil {
		prevNode.Next = nextNode
	}
}

type LRUCache struct {
	Capacity int
	ll       *linkedList
	nodeMap  map[string]*node
}

func NewCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		ll:       newLinkedList(),
		nodeMap:  make(map[string]*node),
	}
}

func (c *LRUCache) Get(key string) (string, error) {
	node, ok := c.nodeMap[key]
	if ok {
		result := node.Value
		c.ll.Remove(node)
		c.ll.AddFront(node)
		return result, nil
	}
	return "", MissingKey{Key: key}
}

func (c *LRUCache) Put(key string, value string) {
	node, ok := c.nodeMap[key]
	if ok {
		node.Value = value
		c.ll.Remove(node)
		c.ll.AddFront(node)
	} else {
		if len(c.nodeMap) == c.Capacity {
			tail := c.ll.Tail
			delete(c.nodeMap, tail.Key)
			c.ll.Remove(tail)
		}
		node := newNode(key, value)
		c.nodeMap[key] = node
		c.ll.AddFront(node)
	}

}
