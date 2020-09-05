package main

import "fmt"

// define peson struct
type Person struct {
	firstName string
	lastName  string
	country   string
	gender    string
	age       int
}

func main() {
	// init person using struct
	person := Person{firstName: "Jeffrin", lastName: "Chittilappilly", country: "India", gender: "M", age: 25}
	// alternative declaration
	// person := Person{"Jeffrin", "Chittilappilly", "India", "M", 26}

	fmt.Println(person)
	fmt.Println(person.firstName)

	person.age++

	fmt.Println(person)
}
