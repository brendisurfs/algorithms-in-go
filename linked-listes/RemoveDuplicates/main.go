package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicates(linkedList *LinkedList) *LinkedList {
	current := linkedList
	store := make(map[int]int)
	for current != nil {
		nextValue := current.Next.Value
		if current.Value == nextValue {
		}

	}
	return nil
}

func internalRemoval(store map[int]int, linkedList *LinkedList) {
	return nil
}

func main() {

}
