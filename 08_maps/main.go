package main

import "fmt"

func main() {
	// // emails := make(map[string]string)
	// emails := make(map[string]string)

	// // Assign maps
	// emails["Jeffrin"] = "jeff@gmail.com"
	// emails["Joseph"] = "joseph@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// declare map and add keyvalue
	emails := map[string]string{
		"Jeffrin": "jeff@gmail.com",
		"Joseph":  "joseph@gmail.com",
		"Mike":    "mike@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(emails["Jeffrin"])
	fmt.Println(len(emails))

	// Delete from a maps
	delete(emails, "Mike")
	fmt.Println(emails)
}
