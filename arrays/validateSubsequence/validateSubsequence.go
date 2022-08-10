package main

import "fmt"

// given a non empty array of ints, write a function that determines
// whether the second array is a subsequence of the first one.

func isValidSub(arr []int, sequence []int) bool {
	arrIdx := 0
	sequenceIdx := 0

	// what we are doing is not checking each item for its value,
	// we are checking that the value passes our check and can move forward.
	for arrIdx < len(arr) && sequenceIdx < len(sequence) {
		if arr[arrIdx] == sequence[sequenceIdx] {
			sequenceIdx += 1
		}
		arrIdx += 1
	}
	return sequenceIdx == len(sequence)
}

func main() {
	testarr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	isValid := isValidSub(testarr, sequence)
	fmt.Println(isValid)
}
