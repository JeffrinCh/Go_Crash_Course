package main

// this is a naive solution.
// here we are comparing each element in array to every other element.
// Time complexity will always be O(n^2) and space complexity O(1)
func InsertionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
