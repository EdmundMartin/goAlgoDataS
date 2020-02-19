package search


func binarySearchHelper(array []int, target, start, end int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if array[middle] == target {
		return middle
	} else if array[middle] < target {
		return binarySearchHelper(array, target, middle+1, end)
	} else {
		return binarySearchHelper(array, target, start, middle-1)
	}
}


func BinarySearch(array []int, target int) int {
	return binarySearchHelper(array, target, 0, len(array)-1)
}