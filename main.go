package main

import "fmt"

func reverseArray(names [4]string) [4]string {
	startIdx := 0
	endIdx := len(names) - 1

	for startIdx < endIdx {
		names[startIdx], names[endIdx] = names[endIdx], names[startIdx]
		startIdx++
		endIdx--
	}

	return names
}

func findMinValue(numbers [3]int) int {
	min := numbers[0]

	for _, v := range numbers {
		if v < min {
			min = v
		}
	}

	return min
}

func checkIsPalindrome(str string) bool {
	isPalindrome := true
	startIdx := 0
	endIdx := len(str) - 1

	for startIdx < endIdx {
		if str[startIdx] != str[endIdx] {
			isPalindrome = false
			break
		}
		startIdx++
		endIdx--
	}

	return isPalindrome
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) prepend(data int) {
	newNode := Node{data: data}

	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
	} else {
		l.head = &newNode
	}
}

func (l *LinkedList) print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) checkIsEmpty() bool {
	return q.front == nil
}

func (q *Queue) peek() (*Node, error) {
	isEmpty := q.checkIsEmpty()
	if isEmpty {
		return nil, fmt.Errorf("Queue is empty")
	} else {
		return q.front, nil
	}
}

func (q *Queue) enqueue(data int) {
	newNode := &Node{data: data}

	if q.front == nil {
		q.front = newNode
	} else {
		q.rear.next = newNode
	}
	q.rear = newNode
}

func (q *Queue) dequeue() (int, error) {
	front := q.front
	isEmpty := q.checkIsEmpty()
	if isEmpty {
		return 0, fmt.Errorf("Queue is empty")
	} else {
		q.front = q.front.next
	}

	return front.data, nil
}

func (q *Queue) print() {
	currentNode := q.front
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	reverseArray([4]string{"John", "Doe", "Smith", "Jim"})
	findMinValue([3]int{3, 2, 1})
	_ = checkIsPalindrome("madam")

	list := LinkedList{}
	list.prepend(1)
	list.prepend(2)
	list.print()

	fmt.Println("----------------------")

	queue := Queue{}
	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)

	front, err := queue.dequeue()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Removing ->", front)
	}

	frontEl, err := queue.peek()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Peeking ->", frontEl.data)
	}

	queue.print()
}
