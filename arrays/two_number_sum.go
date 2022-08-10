package main

import (
	"fmt"
	"sort"
)

func twoNumberSum(arr []int, target int) []int {
	var sumArr []int

	frontPtr := 0
	endPtr := len(arr) - 1

	sorted := arr[:]
	sort.Ints(sorted)

	for frontPtr < endPtr {
		currentSum := sorted[frontPtr] + sorted[endPtr]

		if currentSum == target {
			sumArr = append(sumArr, sorted[frontPtr], sorted[endPtr])
			return sumArr
		} else if currentSum < target {
			frontPtr += 1
		} else if currentSum > target {
			endPtr -= 1
		}
	}

	return sumArr
}

func main() {
	demoArray := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10

	res := twoNumberSum(demoArray, targetSum)
	fmt.Println(res)
}
