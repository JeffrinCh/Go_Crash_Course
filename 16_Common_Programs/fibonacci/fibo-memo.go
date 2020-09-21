package main

func GetNthFib(n int) int {
	var memo = map[int]int{
		1: 0,
		2: 1,
		3: 1,
	}
	return helper(n, memo)
}

func helper(n int, memoize map[int]int) int {

	if val, found := memoize[n]; found {
		return val
	}

	memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
	return memoize[n]
}
