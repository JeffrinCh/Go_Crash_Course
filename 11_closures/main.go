package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// go supports anonymous which support closures
	// we can define a function inline without having to name it.
	sum := adder()

	fmt.Printf("%T\n", sum)
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
