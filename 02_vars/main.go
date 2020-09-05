package main

import "fmt"

// MAIN Types
// string
// int
// bool
// int
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

// using var
// can be delcared outside the function body.
var name string = "Brad"

// should omit the string type as it is inferred from right hand side.
var name1 = "Jeffrin"

var age = 26
var isCool = true

// we cannot short hand syntax outside the function body
// isCooler := false;

func main() {
	fmt.Println(name)
	fmt.Println(name1, age, isCool)
	fmt.Println(name1)
	// gets the type of the variable
	fmt.Printf("%T\n", isCool)

	// shorthand
	isCooler := true
	fmt.Println(isCooler)

	// float64 (by default float64 is inferred)
	size := 5.3
	fmt.Printf("%T\n", size)

	// assigning multiple variables same time
	name2, name3 := "Jeff", "Chittilappilly"
	fmt.Println(name2, name3)

}
