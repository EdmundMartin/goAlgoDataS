package problems

import "sort"

func sum(parts []int) int {
	result := 0
	for _, val := range parts {
		result += val
	}
	return result
}

func ThreeSum(array []int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}
	for i := 0; i < len(array) - 2; i++ {
		if array[i] > 0 {
			return triplets
		}
		if i > 0 && array[i] == array[i-1] {
			continue
		}
		left, right := 0, len(array) - 1
		for left < right {
			current := []int{array[i], array[left], array[right]}
			if sum(current) == 0 {
				triplets = append(triplets, current)
				for left < right && array[left] == array[left+1] {
					left++
				}
				left++
				for left < right && array[right] == array[right-1] {
					right--
				}
				right--
			} else if sum(current) > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return triplets
}