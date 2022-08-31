package main

import "fmt"

func main() {
	tree := &BST{Value: 10, Left: nil, Right: nil}

	doesTreeHave := tree.Contains(5)
	fmt.Printf("%#v\n", doesTreeHave)
}

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	current := tree.Value
	if value > current {
		newTree := tree.Right
		newTree.Insert(value)
	}
	if value < current {
		newTree := tree.Left
		newTree.Insert(value)
	}
	tree.Value = value
	return tree
}

func (tree *BST) Contains(value int) bool {

	if value < tree.Value {
		newTreeLeft := tree.Left
		newTreeLeft.Contains(value)
	}

	if value > tree.Value {
		newTreeRight := tree.Right
		newTreeRight.Contains(value)
	}

	if value == tree.Value {
		return true
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	if tree.Left == nil && tree.Right == nil {
		return tree
	}
	return tree
}
