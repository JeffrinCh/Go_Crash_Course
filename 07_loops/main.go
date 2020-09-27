package main

import (
	"fmt"
)

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 0; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FIzzBuzz Challenge
	for i := 0; i < 100; i++ {
		if i%15+3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
