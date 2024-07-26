package binarytree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}

	t.insert(t.Root, value)
}

func (t *BinaryTree) insert(node *Node, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			t.insert(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			t.insert(node.Right, value)
		}
	}
}

func (t *BinaryTree) Search(value int) bool {
	return t.search(t.Root, value)
}

func (t *BinaryTree) search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return t.search(node.Left, value)
	}

	return t.search(node.Right, value)
}

func (t *BinaryTree) InOrder() []int {
	return t.inOrder(t.Root)
}

func (t *BinaryTree) inOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := t.inOrder(node.Left)
	right := t.inOrder(node.Right)

	return append(append(left, node.Value), right...)
}

func (t *BinaryTree) PreOrder() []int {
	return t.preOrder(t.Root)
}

func (t *BinaryTree) preOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := t.preOrder(node.Left)
	right := t.preOrder(node.Right)

	return append(append([]int{node.Value}, left...), right...)
}

func (t *BinaryTree) PostOrder() []int {
	return t.postOrder(t.Root)
}

func (t *BinaryTree) postOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := t.postOrder(node.Left)
	right := t.postOrder(node.Right)

	return append(append(left, right...), node.Value)
}

func (t *BinaryTree) LevelOrder() []int {
	var result []int
	var queue []*Node

	if t.Root != nil {
		queue = append(queue, t.Root)
	}

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
