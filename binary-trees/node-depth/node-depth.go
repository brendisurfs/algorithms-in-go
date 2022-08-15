package main

import "fmt"

// write a function that takes in a binary tree
// and returns the sum of its node depths.

// - node depth: distance between a node and the tree's root.

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func CalcNodeDepths(root *BinaryTree) int {
	sum := 0
	recursiveCalc(root, 0, &sum)
	return sum
}

func recursiveCalc(node *BinaryTree, depth int, sum *int) {
	if node == nil {
		return
	}

	*sum += depth
	recursiveCalc(node.Left, depth+1, sum)
	recursiveCalc(node.Right, depth+1, sum)
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
	depth := CalcNodeDepths(root)
	fmt.Println(depth)
}
