package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// takes a binary tree and returns a list of branch sums
// ordered from leftmost branch sum to rightmost branch sum.
func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	calcBranchSums(root, 0, &sums)
	return sums
}

// calculate using the left and right node recursively.
func calcBranchSums(node *BinaryTree, movingSum int, sums *[]int) {
	if node == nil {
		return
	}

	movingSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, movingSum)
		return
	}
	calcBranchSums(node.Left, movingSum, sums)
	calcBranchSums(node.Right, movingSum, sums)

}

func main() {
	firstRight := &BinaryTree{
		Value: 3,
		Left: &BinaryTree{
			Value: 6,
		},
		Right: &BinaryTree{
			Value: 7,
		},
	}

	firstLeft := &BinaryTree{
		Value: 2,
		Left: &BinaryTree{
			Value: 4,
			Left: &BinaryTree{
				Value: 8,
			},
			Right: &BinaryTree{
				Value: 9,
			},
		},
		Right: &BinaryTree{
			Value: 4,
			Left: &BinaryTree{
				Value: 10,
			},
		},
	}

	var root = &BinaryTree{
		Value: 1,
		Left:  firstLeft,
		Right: firstRight,
	}

	sums := BranchSums(root)
	fmt.Println(sums)
}
