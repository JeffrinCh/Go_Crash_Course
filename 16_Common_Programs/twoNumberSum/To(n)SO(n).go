package main

//To(n)SO(n)
func TwoNumberSum(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		val := target - num
		if _, found := nums[val]; found {
			return []int{num, val}
		} else {
			nums[num] = true
		}
	}
	return []int{}
}
