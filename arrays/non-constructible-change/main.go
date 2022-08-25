package main

import (
	"fmt"
	"sort"
)

// given an array of positive integers representing the values of coins in your
// possession, write a function that returns the minimum amount of change
// that you cannot create.
// you CAN have multiple of the same coin.

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	NonConstructibleChange(coins)
}

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	change := 0

	for i := range coins {
		change += coins[i]
		fmt.Println(change)
	}

	return 0
}
