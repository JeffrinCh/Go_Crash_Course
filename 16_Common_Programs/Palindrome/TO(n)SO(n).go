package main
//import "fmt"

// Time Complexity O(n) Space o(n)
func IsPalindrome(str string) bool {
	length := len(str)
	for i, r := range str {
		
		if int(r) != int(str[length - i - 1]) {
			return false
		}
	}
	return true
}
