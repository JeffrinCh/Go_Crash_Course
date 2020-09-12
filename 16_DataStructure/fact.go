package main

import (
	"fmt"
)

func fact(i int) int {
	if i >= 1 {
		return i * fact(i-1)
	} else {
		return 1
	}
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
	fmt.Print(fact(i))
}
