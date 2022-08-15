package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

func (node *Node) DepthFirstSearch(array []int) []string {
	if node == nil {
		return nil
	}
	return nil
}

func main() {
	testNode := Node{
		Name: "A",
		Children: []*Node{
			&Node{
				Name: "B",
				Children: []*Node{
					&Node{Name: "E"},
					&Node{
						Name:     "F",
						Children: []*Node{&Node{Name: "I"}, &Node{Name: "J"}},
					},
				},
			},
			&Node{Name: "C"},
			&Node{
				Name: "D",
				Children: []*Node{
					&Node{Name: "G", Children: []*Node{&Node{Name: "K"}}},
					&Node{Name: "H"},
				},
			},
		},
	}

	fmt.Printf("%#v\n", testNode)
}
