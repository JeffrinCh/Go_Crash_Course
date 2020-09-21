package main

import (
	"fmt"
)

func GetNthFib(n int) int {

	fibArr := []int{0, 1}
	for counter := 3; counter <= n; counter++ {
		nextFib := fibArr[0] + fibArr[1]
		if counter == n {
			fmt.Print(nextFib)
			return nextFib
		}
		fibArr[0] = fibArr[1]
		fibArr[1] = nextFib
	}
	if n == 2 {
		return 1
	}
	return 0
}
