package avltree

import (
	"fmt"

	"github.com/mounis-bhat/dsa-go/pkg/utils"
)

type Node struct {
	Value  int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func height(n *Node) int {
	if n == nil {
		return -1
	}
	return n.Height
}

func balanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func rotateRight(currentNode *Node) *Node {
	newNode := currentNode.Left
	T2 := newNode.Right

	newNode.Right = currentNode
	currentNode.Left = T2

	currentNode.Height = 1 + utils.Max(height(currentNode.Left), height(currentNode.Right))
	newNode.Height = 1 + utils.Max(height(newNode.Left), height(newNode.Right))

	return newNode
}

func rotateLeft(currentNode *Node) *Node {
	newNode := currentNode.Right
	T2 := newNode.Left // which is the old newNode.Left

	newNode.Left = currentNode
	currentNode.Right = T2

	currentNode.Height = 1 + utils.Max(height(currentNode.Left), height(currentNode.Right))
	newNode.Height = 1 + utils.Max(height(newNode.Left), height(newNode.Right))

	return newNode
}

func (t *AVLTree) Insert(value int) {
	t.Root = insert(t.Root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value, Height: 0}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + utils.Max(height(node.Left), height(node.Right))

	balance := balanceFactor(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Value {
		return rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Value {
		return rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Value {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Value {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func (t *AVLTree) InOrder() []int {
	return inOrder(t.Root)
}

func inOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	return append(append(inOrder(node.Left), node.Value), inOrder(node.Right)...)
}

func (t *AVLTree) PreOrder() []int {
	return preOrder(t.Root)
}

func preOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	return append(append([]int{node.Value}, preOrder(node.Left)...), preOrder(node.Right)...)
}

func (t *AVLTree) PostOrder() []int {
	return postOrder(t.Root)
}

func postOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	return append(append(postOrder(node.Left), postOrder(node.Right)...), node.Value)
}

func (t *AVLTree) LevelOrder() []int {
	var result []int
	if t.Root == nil {
		return result
	}

	queue := []*Node{t.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func visualizeNode(node *Node, prefix string, isLeft bool) string {
	if node == nil {
		return ""
	}

	output := ""

	if node.Right != nil {
		output += visualizeNode(node.Right, prefix+(map[bool]string{true: "│   ", false: "    "})[isLeft], false)
	}

	output += prefix + (map[bool]string{true: "└── ", false: "┌── "})[isLeft] + fmt.Sprintf("%d\n", node.Value)

	if node.Left != nil {
		output += visualizeNode(node.Left, prefix+(map[bool]string{true: "    ", false: "│   "})[isLeft], true)
	}

	return output
}

func (t *AVLTree) Visualize() {
	fmt.Println(visualizeNode(t.Root, "", true))
}
