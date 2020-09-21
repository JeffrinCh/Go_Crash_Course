package main

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Time Complexity O(log) worst case O(d) Space complexity O(n)
func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(tree.Value, target)
}

func (tree *BST) findClosestValue(closestValue int, target int) int {
	if int(math.Abs(float64(closestValue-target))) > int(math.Abs(float64(tree.Value-target))) {
		closestValue = tree.Value
	}

	if int(math.Abs(float64(tree.Value-target))) == 0 {
		return tree.Value
	}
	if target < tree.Value {
		if tree.Left != nil {
			return tree.Left.findClosestValue(closestValue, target)
		}
	} else if target > tree.Value {
		if tree.Right != nil {
			return tree.Right.findClosestValue(closestValue, target)
		}
	}
	return closestValue
}
