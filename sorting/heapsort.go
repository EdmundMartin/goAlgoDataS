package sorting

type heap struct {
	values []int
}

func newHeap() *heap {
	return &heap{
		values: []int{},
	}
}


func (h *heap) isLeaf(idx int) bool {
	if idx >= len(h.values) / 2 && idx <= len(h.values) {
		return true
	}
	return false
}

func (h *heap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *heap) leftChild(idx int) int {
	return (2 * idx) + 1
}

func (h *heap) rightChild(idx int) int {
	return (2 * idx) + 2
}

func (h *heap) insert(value int) {
	h.values = append(h.values, value)
	h.upHeapify(len(h.values)-1)
}

func (h *heap) swap(first, second int) {
	temp := h.values[first]
	h.values[first] = h.values[second]
	h.values[second] = temp
}

func (h *heap) upHeapify(idx int) {
	for h.values[idx] < h.values[h.parent(idx)] {
		h.swap(idx, h.parent(idx))
	}
}

func (h *heap) downHeapify(current int) {
	if h.isLeaf(current) {
		return
	}
	smallest := current
	leftChild := h.leftChild(current)
	rightChild := h.rightChild(current)
	currentSize := len(h.values)
	if leftChild < currentSize && h.values[leftChild] < h.values[smallest] {
		smallest = leftChild
	}
	if rightChild < currentSize && h.values[rightChild] < h.values[smallest] {
		smallest = rightChild
	}
	if smallest != current {
		h.swap(current, smallest)
		h.downHeapify(smallest)
	}
	return
}


func (h *heap) Remove() int {
	top := h.values[0]
	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	h.downHeapify(0)
	return top
}

func (h *heap) buildMinHeap() {
	for idx := ((len(h.values) / 2) - 1); idx >= 0; idx-- {
		h.downHeapify(idx)
	}
}

func HeapSort(array []int) []int {
	h := newHeap()
	for _, val := range array {
		h.insert(val)
	}
	sorted := []int{}
	h.buildMinHeap()
	for len(h.values) > 0 {
		top := h.Remove()
		sorted = append(sorted, top)
	}
	return sorted
}