package main

import (
	"fmt"
)

func main() {
	// // Arrays
	// var fruitArray [2]string

	// // in go arrays have to be a fixed length
	// // slice is basically an array which doesn;t have fixed type

	// // Assign Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// declare and assign
	fruitArray := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])

	// slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)

	// get lenght of the array
	fmt.Println(len(fruitSlice))

	// range of array
	fmt.Println(fruitSlice[1:3])
}
