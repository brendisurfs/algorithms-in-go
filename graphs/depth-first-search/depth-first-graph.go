package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

func (node *Node) DepthFirstSearch(array []string) []string {
	array = append(array, node.Name)
	for _, child := range node.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
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

	store := []string{}
	res := testNode.DepthFirstSearch(store)
	fmt.Println(res)
}
