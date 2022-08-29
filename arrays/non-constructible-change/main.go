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
	ncc := NonConstructibleChange(coins)
	fmt.Println(ncc)
}

func NonConstructibleChange(coins []int) int {
	change := 0
	sort.Ints(coins)

	for _, coin := range coins {
		fmt.Printf("coin: %d, change: %d\n", coin, change)
		if coin > change+1 {
			return (change + 1)
		}
		change += coin
	}
	return 0
}
