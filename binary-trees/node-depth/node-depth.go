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
	depthList := []int{}
	counter := 0
	internalCalc(root, counter, &depthList)

	sum := 0
	for _, v := range depthList {
		sum += v
	}
	return sum
}

func internalCalc(node *BinaryTree, counter int, depthList *[]int) {

	if node == nil {
		return
	}
	counter += 1

	if node.Left == nil && node.Right == nil {
		*depthList = append(*depthList, counter)
		return
	}

	internalCalc(node.Left, counter, depthList)
	internalCalc(node.Right, counter, depthList)
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
