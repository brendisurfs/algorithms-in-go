package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (t *BST) FindClosestValue(target int) int {
	return t.findClosestValue(target, t.Value)
}

func (t *BST) findClosestValue(target, closest int) int {
	current := t

	for current != nil {
		if DiffAB(target, closest) > DiffAB(target, current.Value) {
			closest = current.Value
		}
		if target < current.Value {
			current = current.Left
		} else if target > current.Value {
			current = current.Right
		} else {
			break
		}
	}
	return closest
}

func DiffAB(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {

	testTree := BST{
		// root
		Value: 10,

		// first level left
		Left: &BST{
			Value: 5, Left: &BST{
				Value: 2, Left: &BST{
					Value: 1,
				},
			},
			Right: &BST{
				Value: 5,
			},
		},

		// first level right
		Right: &BST{
			Value: 15,
			Right: &BST{
				Value: 22,
			},
			Left: &BST{
				Value: 13,
				Right: &BST{
					Value: 14,
				},
			},
		},
	}

	testTarget := 12

	closestValue := testTree.FindClosestValue(testTarget)
	fmt.Println(closestValue)
}
