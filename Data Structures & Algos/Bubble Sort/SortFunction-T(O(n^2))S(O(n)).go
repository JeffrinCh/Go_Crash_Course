package main

// TO(n^2) S O(n)
func BubbleSort(array []int) []int {
	swap := false
	for i := 0; i < len(array) -1 ; i++ {
		if array[i] > array[i + 1] {
			swap = true
			array[i], array[i + 1] = array[i + 1], array[i]
		}
	}
	if (swap) {
		return BubbleSort(array)
	}
	return array
}
