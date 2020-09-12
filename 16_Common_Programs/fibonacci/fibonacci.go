package main

import (
	"fmt"
)

func GetNthFib(n int) int {
	if n > 3 {
		return GetNthFib(n-1) + GetNthFib(n-2)
	} else if n == 1 {
		return 0
	}
	return 1
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Print(GetNthFib(n))
}
