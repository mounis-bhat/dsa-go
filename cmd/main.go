package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Prepend(data int) {
	newNode := Node{Data: data}

	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = &newNode
	}

}

func (ll *LinkedList) Print() {
	currentNode := ll.Head

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

func (ll *LinkedList) Append(data int) {
	newNode := Node{Data: data}
	currentNode := ll.Head

	if ll.Head == nil {
		ll.Head = &newNode
		return
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = &newNode

}

func (ll *LinkedList) GetLength() int {
	length := 0
	currentNode := ll.Head

	for currentNode != nil {
		currentNode = currentNode.Next
		length++
	}

	return length
}

func (ll *LinkedList) InsertAt(position, data int) {
	length := ll.GetLength()

	if position > length || position < 0 {
		fmt.Println("Out of bounds")
		return
	}

	if position == 0 {
		ll.Prepend(data)
		return
	}

	if position == length {
		ll.Append(data)
		return
	}

	currentNode := ll.Head
	newNode := Node{Data: data}
	i := 0
	for i < position-1 {
		currentNode = currentNode.Next
		i++
	}

	newNode.Next = currentNode.Next
	currentNode.Next = &newNode
}

func (ll *LinkedList) DeleteAt(position int) {
	length := ll.GetLength()

	if position >= length || position < 0 {
		fmt.Println("Out of bounds")
		return
	}
	currentNode := ll.Head

	if position == 0 {
		ll.Head = currentNode.Next
		return
	}

	i := 0

	for i < position-1 {
		currentNode = currentNode.Next
		i++
	}

	currentNode.Next = currentNode.Next.Next
}

func main() {
	myList := NewLinkedList()
	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	myList.Append(4)
	myList.Append(5)
	myList.InsertAt(2, 25)
	myList.DeleteAt(2)
	myList.Print()

}
