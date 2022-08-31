package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	tree := &BST{Value: 10}

	tree.Insert(5)
	contains := tree.Contains(12)
	removedTree := tree.Remove(5).ToPretty()

	fmt.Printf("%#v\n", contains)
	fmt.Printf("%+v\n", removedTree)
}

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Time: O(Log(N)) | Space: O(1)
func (tree *BST) Insert(value int) *BST {
	current := tree
	for {
		if value > current.Value {
			if current.Left == nil {
				current.Left = &BST{Value: value}
				break
			} else {
				current = current.Left
			}
		} else if current.Right == nil {
			current.Right = &BST{Value: value}
			break
		} else {
			current = current.Right
		}
	}
	return tree
}

// worst case Time: O(n) | Space: O(1)
func (tree *BST) Contains(value int) bool {
	current := tree
	for current != nil {
		if value > current.Value {
			current = current.Right
		} else if value < current.Value {
			current = current.Left
		} else {
			return true
		}
	}
	return false
}

// removal is a two step process.
// 1. find the node
// 2. Remove the node.
func (tree *BST) Remove(value int, parentNode *BST) *BST {

	current := tree

	for current != nil {
		if value > current.Value {
			parentNode = current
			current = current.Right
		} else if value < current.Value {
			parentNode = current
			current = current.Left
		} else {
			if current.Left != nil && current.Right != nil {
				// need to get rid of ALL the nodes after the node,
				// not just that nodes value.
				current.Value = current.Right.getMinValue()
				current.Right.Remove(current.Value, current)
			}
		}
	}
	return tree
}

func (tree *BST) getMinValue() int {
	fmt.Println("not implemented")
	return -1
}

func (tree *BST) ToPretty() string {
	parsed, err := json.MarshalIndent(tree, "", " ")
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(parsed)
}
