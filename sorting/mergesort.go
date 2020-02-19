package sorting

func merge(array []int, left, middle, right int) {
	n1 := middle - left + 1
	n2 := right - middle

	tmpLeft := make([]int, n1)
	tmpRight := make([]int, n2)

	for i := 0; i < n1; i++ {
		tmpLeft[i] = array[left+i]
	}
	for j := 0; j < n2; j++ {
		tmpRight[j] = array[middle+1+j]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if tmpLeft[i] <= tmpRight[j] {
			array[k] = tmpLeft[i]
			i++
		} else {
			array[k] = tmpRight[j]
			j++
		}
		k++
	}

	for i < n1 {
		array[k] = tmpLeft[i]
		i++
		k++
	}

	for j < n2 {
		array[k] = tmpRight[j]
		j++
		k++
	}
}

func MergeSort(array []int, start, end int) {
	if start < end {
		middle := (start + end) / 2
		MergeSort(array, start, middle)
		MergeSort(array, middle+1, end)

		merge(array, start, middle, end)
	}
}
