package main

// O(n^2) Time Complexity and O(n) space complexity
func TwoNumberSum(array []int, target int) []int {
	len := len(array)
	for i, value := range array {
		for j := i + 1; j < len; j++ {
			if (value + array[j]) == target {
				return []int{value, array[j]}
			}
		}
	}
	return []int{}
}