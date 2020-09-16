package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(tree.Value, target)
}

func (tree *BST) findClosestValue(closestValue int, target int) int {
	currentNode := tree
	for currentNode != nil {
		if int(math.Abs(float64(target-closestValue))) > int(math.Abs(float64(target-currentNode.Value))) {
			closestValue = currentNode.Value
			fmt.Println(closestValue)
		}
		if int(math.Abs(float64(target-currentNode.Value))) == 0 {
			closestValue = currentNode.Value
		}
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return closestValue
}
