package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) AddNode(value int) *LinkedList {
	l.Next = &LinkedList{Value: value}
	return l.Next
}

func RemoveDuplicates(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList

	for currentNode != nil {
		// this will allow us to check against nodes.
		nextNode := currentNode.Next
		for nextNode != nil && nextNode.Value == currentNode.Value {
			nextNode = nextNode.Next
		}
		currentNode.Next = nextNode
		currentNode = nextNode
	}

	return linkedList
}

func main() {
	testList := LinkedList{Value: 1}
	testList.AddNode(3).AddNode(4).AddNode(4).AddNode(5).AddNode(6).AddNode(6)
	temp := RemoveDuplicates(&testList)
	fbytes, err := json.MarshalIndent(temp, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fbytes))
}
