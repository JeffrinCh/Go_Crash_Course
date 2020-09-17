package main

//import "fmt"

func IsPalindrome(str string) bool {
	return IsPalindromeHelper(str, 0, false)
}

func IsPalindromeHelper(str string, counter int, isPalindrome bool) bool {
	if len(str) == counter {
		return true
	}
	if int(str[counter]) == int(str[len(str)-counter-1]) {
		counter++
		isPalindrome = IsPalindromeHelper(str, counter, isPalindrome)
	} else {
		return false
	}
	return isPalindrome
}
