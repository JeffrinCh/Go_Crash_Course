package main

import (
	"fmt"
	"strconv"
)

// define peson struct
type Person struct {
	firstName, lastName, country, gender string
	age                                  int
}

// greeting method(value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + "  and I am " + strconv.Itoa(p.age)
}

// hasBirthday method(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spuseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spuseLastName
	}
}

func main() {

	// we have two kinds of methods value receivers and pointer receivers

	// value receivers are for method like that do calculations.
	// They dont actually change anything

	// pointer receivers are for when ur actually changing something
	// init person using struct
	person := Person{firstName: "Arthi", lastName: "WIllaims", country: "India", gender: "F", age: 25}

	// alternative declaration
	person1 := Person{"Bob", "Jhonson", "New York", "M", 26}

	person.hasBirthday()
	fmt.Println(person.greet())
	person.getMarried("Jeff")
	person1.getMarried("Thomson")
	fmt.Println(person.greet())
}
