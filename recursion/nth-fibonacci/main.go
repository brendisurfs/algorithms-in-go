package main

import (
	"fmt"
)

func main() {
	store := make(map[int]int)
	store[0] = 0
	store[1] = 1

	nth := GetNthFib(9, store)
	fmt.Println(nth)

	iterNth := IterNthFib(9)
	fmt.Println(iterNth)
}

// O(n) time | O(n) space
func GetNthFib(n int, store map[int]int) int {
	_, found := store[n]
	if found {
		return store[n]
	} else {
		store[n] = GetNthFib(n-1, store) + GetNthFib(n-2, store)
		return store[n]
	}
}

// O(n) time | O(1) space
func IterNthFib(n int) int {

	lastTwo := []int{0, 1}
	counter := 2

	for counter <= n {
		next := lastTwo[0] + lastTwo[1]
		lastTwo[0] = lastTwo[1]
		lastTwo[1] = next
		counter += 1
	}
	if n > 1 {
		return lastTwo[1]
	} else {
		return lastTwo[0]
	}
}
