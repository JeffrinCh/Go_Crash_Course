package main
import (
	"sort"
)

// TO(nlog(n)SO(1)
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array);
	left := 0;
	right := len(array)-1
	for i := 0; i < len(array)-1; i ++ {
		if (array[left] + array[right]  >  target) {
			right -= 1
		} else if array[left] + array[right]  < target {
			left += 1
		} else {
			return []int{array[left], array[right]}
		}
	}
	return []int{}
}
