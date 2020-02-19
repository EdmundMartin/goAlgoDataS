package sorting


func QuickSort(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	if len(array) == 1 {
		return array
	}
	pivot := array[0]
	smaller := []int{}
	larger := []int{}
	i := 1
	for i < len(array) {
		if array[i] < pivot {
			smaller = append(smaller, array[i])
		} else {
			larger = append(larger, array[i])
		}
		i++
	}
	smaller = QuickSort(smaller)
	larger = QuickSort(larger)
	smaller = append(smaller, pivot)
	smaller = append(smaller, larger...)
	return smaller
}