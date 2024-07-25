package main

import (
	"fmt"
	"math"
)

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
	}

	l.head = &newNode
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

type BTNode struct {
	data  int
	left  *BTNode
	right *BTNode
}

type BinaryTree struct {
	root *BTNode
}

func (b *BinaryTree) depthFirstTraverse() {
	if b.root == nil {
		return
	}
	stack := []*BTNode{b.root}
	for len(stack) > 0 {
		lastIdx := len(stack) - 1

		current := stack[lastIdx]

		// removing last element, because lifo
		stack = stack[:lastIdx]

		fmt.Println(current.data)

		if current.right != nil {
			stack = append(stack, current.right)
		}

		if current.left != nil {
			stack = append(stack, current.left)
		}
	}
}

func preOrder(node *BTNode) {
	if node != nil {
		fmt.Println(node.data)
		preOrder(node.left)
		preOrder(node.right)
	}
}

func inOrder(node *BTNode) {
	if node != nil {
		preOrder(node.left)
		fmt.Println(node.data)
		preOrder(node.right)
	}
}

func postOrder(node *BTNode) {
	if node != nil {
		preOrder(node.left)
		preOrder(node.right)
		fmt.Println(node.data)
	}
}

func (b *BinaryTree) recursiveDFT() {
	fmt.Println("Pre order --------")
	preOrder(b.root)
	fmt.Println("In order--------")
	inOrder(b.root)
	fmt.Println("Post order--------")
	postOrder(b.root)
}

func (b *BinaryTree) breadthFirst() {
	if b.root == nil {
		return
	}

	queue := []*BTNode{b.root}

	for len(queue) > 0 {
		current := queue[0]
		// remove the first element from the queue, because fifo
		queue = queue[1:]

		fmt.Println(current.data)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func insertNode(node *BTNode, data int) {
	if data < node.data {
		if node.left == nil {
			node.left = &BTNode{data: data}
		} else {
			insertNode(node.left, data)
		}
	} else {
		if node.right == nil {
			node.right = &BTNode{data: data}
		} else {
			insertNode(node.right, data)
		}
	}
}

func (b *BinaryTree) insert(data int) {
	if b.root == nil {
		b.root = &BTNode{data: data}
	} else {
		insertNode(b.root, data)
	}
}

func sum(root *BTNode) int {
	totalSum := 0
	stack := []*BTNode{root}
	for len(stack) > 0 {
		lastIdx := len(stack) - 1

		current := stack[lastIdx]

		stack = stack[:lastIdx]

		totalSum += current.data

		if current.right != nil {
			stack = append(stack, current.right)
		}

		if current.left != nil {
			stack = append(stack, current.left)
		}
	}
	return totalSum
}

func (b *BinaryTree) calculateSum() int {
	if b.root == nil {
		return 0
	}
	return sum(b.root)
}

func recursiveSum(node *BTNode) int {
	if node == nil {
		return 0
	}
	return node.data + recursiveSum(node.left) + recursiveSum(node.right)
}

func (b *BinaryTree) calculateRecursiveSum() int {
	if b.root == nil {
		return 0
	}
	return recursiveSum(b.root)
}

// func performIns(node *BTNode, data int) {
// 	if data < node.data {
// 		if node.left == nil {
// 			node.left = &BTNode{data: data}
// 		} else {
// 			performIns(node.left, data)
// 		}
// 	} else {
// 		if node.right == nil {

// 			node.right = &BTNode{data: data}
// 		} else {
// 			performIns(node.right, data)
// 		}
// 	}
// }

// func (b *BinaryTree) insertionPrac(data int) {
// 	if b.root == nil {
// 		b.root = &BTNode{data: data}
// 	}

// 	performIns(b.root, data)
// }

func compareNodes(a, b, c int) int {
	switch {
	case a < b && a < c:
		return a
	case b < c && b < a:
		return b
	default:
		return c
	}
}

func findMin(node *BTNode) int {
	if node == nil {
		return math.MaxInt
	}
	return compareNodes(node.data, findMin(node.left), findMin(node.right))
}

func maxRootToLeafPathSum(node *BTNode) int {
	if node == nil {
		return math.MinInt
	}

	if node.left == nil && node.right == nil {
		return node.data
	}

	left := maxRootToLeafPathSum(node.left)
	right := maxRootToLeafPathSum(node.right)

	if left > right {
		return node.data + left
	}

	return node.data + right
}

func main() {
	reverseArray([4]string{"John", "Doe", "Smith", "Jim"})
	findMinValue([3]int{3, 2, 1})
	_ = checkIsPalindrome("madam")

	list := LinkedList{}
	list.prepend(1)
	list.prepend(2)
	list.print()

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

	tree := BinaryTree{}

	tree.insert(40)
	tree.insert(30)
	tree.insert(35)
	tree.insert(20)
	tree.insert(25)
	tree.insert(60)
	tree.insert(80)
	tree.insert(70)
	tree.insert(75)
	tree.insert(65)

	fmt.Println("DFT Stack----------")
	tree.depthFirstTraverse()
	tree.recursiveDFT()
	fmt.Println("BFT Queue----------")
	tree.breadthFirst()

	sum := tree.calculateSum()
	fmt.Println("Sum----------")
	fmt.Println(sum)
	fmt.Println("Rec Sum----------")
	recSum := tree.calculateRecursiveSum()
	fmt.Println(recSum)

	fmt.Println("Min----------")
	min := findMin(tree.root)
	fmt.Println(min)

	fmt.Println("Max Root to Leaf path sum----------")
	maxRTLPSum := maxRootToLeafPathSum(tree.root)
	fmt.Println(maxRTLPSum)
}
