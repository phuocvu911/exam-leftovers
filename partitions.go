package main

import (
	"fmt"
)

func Partitions(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	current := []int{}
	result := [][]int{}

	var backtrack func(remainder, max int) // var backtrack is an anonymous func
	backtrack = func(remainder, max int) {
		if remainder == 0 {
			subres := make([]int, len(current)) // have to copy from current ecause we change current all the time  and if we append current into result we onlyh have all of the last arrays since slice is a pointer to an underlying array.
			copy(subres, current)               // have to use copy function
			result = append(result, subres)
			return
		}

		for i := min(remainder, max); i >= 1; i-- { // use min, so that the items got added later always less than or equal to the first item in the subArr(partitions)
			current = append(current, i)
			backtrack(remainder-i, i)
			current = current[:len(current)-1] // reset current after used.
		}
	}
	backtrack(n, n)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Partitions(5))

}
