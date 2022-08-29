package main

import (
	"fmt"
)

func main() {
	nth := GetNthFib(9)
	fmt.Println(nth)
}

func GetNthFib(n int) int {
	store := make(map[int]int)

	intSlice := []int{0, 1}

	for i := 0; i <= n; i++ {
		next := intSlice[0] + intSlice[1]
		store[i] = next
		intSlice[0] = intSlice[1]
		intSlice[1] = next
	}

	return store[n]
}
