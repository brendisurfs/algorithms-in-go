package main

import "fmt"

// given an array of positive integers representing the values of coins in your
// possession, write a function that returns the minimum amount of change
// that you cannot create.
// you CAN have multiple of the same coin.

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	NonConstructibleChange(coins)
}

func NonConstructibleChange(coins []int) int {
	for i := range coins {
		fmt.Println(i)
	}
	return 0
}
