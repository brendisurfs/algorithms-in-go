package main

import (
	"fmt"
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
	nextNode := currentNode.Next

	// tempNode := currentNode.Next.Next

	for currentNode.Next != nil {
		fmt.Printf("%+v\n", currentNode)
		if currentNode.Value == nextNode.Value {
			currentNode = currentNode.Next
		}
		currentNode = currentNode.Next
	}
	return nil
}

func main() {
	testList := LinkedList{Value: 1}
	testList.AddNode(3).AddNode(4).AddNode(4).AddNode(5).AddNode(6).AddNode(6)
	// fbytes, err := json.MarshalIndent(testList, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	RemoveDuplicates(&testList)
	// fmt.Println(string(fbytes))
}
