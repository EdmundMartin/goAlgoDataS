package problems

import "sort"

func TwoSum(array []int, target int) [][]int {
	sort.Ints(array)
	found := [][]int{}
	other := make(map[int]bool)
	idx := 0
	for idx < len(array) {
		remaining := target - array[idx]
		_, ok := other[remaining]
		if ok {
			found = append(found, []int{remaining, array[idx]})
			for idx < len(array) - 1 && array[idx] == array[idx+1] {
				idx++
			}
		}
		other[array[idx]] = true
		idx++
	}
	return found
}
