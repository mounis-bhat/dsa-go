package linkedlist

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
	newNode := &Node{Data: data}
	if ll.Head != nil {
		newNode.Next = ll.Head
	}
	ll.Head = newNode
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) Display() []int {
	var elements []int
	current := ll.Head
	for current != nil {
		elements = append(elements, current.Data)
		current = current.Next
	}
	return elements
}

func (ll *LinkedList) Insert(data int, index int) {
	if index == 0 {
		ll.Prepend(data)
		return
	}
	newNode := &Node{Data: data}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	if current == nil {
		return
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func (ll *LinkedList) Delete(data int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (ll *LinkedList) Reverse() {
	var previous *Node
	var next *Node
	current := ll.Head
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	ll.Head = previous
}

func (ll *LinkedList) Get(index int) int {
	current := ll.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return -1
		}
		current = current.Next
	}
	if current == nil {
		return -1
	}
	return current.Data
}

func (ll *LinkedList) Search(data int) int {
	current := ll.Head
	index := 0
	for current != nil {
		if current.Data == data {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

func (ll *LinkedList) Size() int {
	current := ll.Head
	size := 0
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

func (ll *LinkedList) Clear() {
	ll.Head = nil
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *LinkedList) Shift() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}
	data := ll.Head.Data
	ll.Head = ll.Head.Next
	return data, true
}

func (ll *LinkedList) Pop() (int, bool) {
	current := ll.Head
	if current == nil {
		return 0, false
	}
	if current.Next == nil {
		ll.Head = nil
		return current.Data, true
	}
	var next *Node
	for current != nil {
		next = current.Next
		if next.Next == nil {
			current.Next = nil
			break
		}
		current = current.Next
	}
	return next.Data, true
}
