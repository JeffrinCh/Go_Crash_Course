package main

import (
	"fmt"
)

func greetings(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Println(greetings("Jolly"))
}
