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
	return l
}

func RemoveDuplicates(linkedList *LinkedList) *LinkedList {
	current := linkedList

	for current.Next != nil {
		next := current.Next
		if current.Value == next.Value {
			next = next.Next
			current = next
		}

	}
	return nil
}

func main() {
	testList := LinkedList{Value: 1}
	testList.AddNode(3).AddNode(4).AddNode(4).AddNode(5).AddNode(6).AddNode(6)
	fbytes, err := json.MarshalIndent(testList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fbytes))
}
