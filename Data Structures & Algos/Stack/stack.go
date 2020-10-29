package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// return top element of the stack, return false if empty
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	idx := len(*s) - 1
	topElement := (*s)[idx]
	*s = (*s)[:idx]
	return topElement, true
}

func (s *Stack) Push(str string) bool {
	*s = append(*s, str)
	return true
}

func main() {
	var stack Stack
	stack.Push("Helllo")
	stack.Push("Jeffrin")
	stack.Push("This is the top element")
	fmt.Println(stack.Pop())
}
