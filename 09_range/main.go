package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop thru ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using ids
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	// Range with maps
	emails := map[string]string{
		"Jeffrin": "jeff@gmail.com",
		"bob":     "bob@gmail.com",
		"mike":    "mike@gmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// printing just keys
	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
