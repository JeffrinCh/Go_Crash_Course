package main

import (
	"fmt"
)

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	// type of pointer
	fmt.Printf("%T\n", b) // *int (* represents a pointer *int and int are two different data types)

	// Use * to read val from address
	fmt.Println(*b)

	// Change val with pointers
	*b = 10
	fmt.Println(*b)

	// reason to use pointer because you might have to pass a lot of data stored at an address
	// and if we chose to pass the address instead of the data itself it can increase performance
	// simply make things faster

	// everything in go is pased by value
}
