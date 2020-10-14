package main

type SpecialArray []interface{}

// O(n) time | O)d_ space -where n is he total number of elements in the array. uincluding sub-elements
// and d  is the greatest depth of "sepcial arrays " in the arrays
func ProductSum(array SpecialArray) int {
	return helper(array, 1)
}

func helper(array SpecialArray, depth int) int {
	sum := 0
	for _, el := range array {
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, depth+1)
		} else if cast, ok := el.(int); ok {
			sum += cast
		}
	}
	return sum * depth
}
