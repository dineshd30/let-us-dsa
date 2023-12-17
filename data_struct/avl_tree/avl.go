package main

import "fmt"

const SPACE string = "     "

// AVL tree structure i.e. self balancing bst
type AVL struct {
	root *Node
}

// AVL node structure
type Node struct {
	value int
	left  *Node
	right *Node
}

// New AVL tree
func NewAVL(arr []int) *AVL {
	avl := &AVL{}
	for _, v := range arr {
		avl.Insert(v)
	}
	return avl
}

// New AVL node
func NewNode(v int) *Node {
	return &Node{value: v}
}

// Wrapper insert into AVL
func (a *AVL) Insert(v int) {
	fmt.Printf("Inserting %d\n", v)
	a.root = a.root.Insert(v)
	a.root.Print("")
}

// Internal insert into AVL
func (node *Node) Insert(v int) *Node {
	// Base condition
	if node == nil {
		node := NewNode(v)
		return node
	}

	if v > node.value {
		node.right = node.right.Insert(v)
	} else {
		node.left = node.left.Insert(v)
	}

	return node.Balance(v)
}

// Balance the given node of AVL tree
func (node *Node) Balance(v int) *Node {
	// Get the balance factor for each node
	bf := node.GetBalanceFactor()

	if bf == -1 || bf == 0 || bf == 1 {
		return node
	} else {
		fmt.Printf("bf of node %v is %d, balancing the avl tree\n", node, bf)
	}

	if bf > 1 && v < node.left.value { // Left Left Imbalance; do right rotation
		node = node.RotateRight()
	} else if bf < -1 && v > node.right.value { // Right Right Imbalance; do left rotation
		node = node.RotateLeft()
	} else if bf > 1 && v > node.left.value { // Left Right Imbalance; do left right rotation
		node.left = node.left.RotateLeft()
		node = node.RotateRight()
	} else if bf < -1 && v < node.right.value { // Right Left Imbalance; do right left rotation
		node.right = node.right.RotateRight()
		node = node.RotateLeft()
	}

	return node
}

// Rotate left
func (root *Node) RotateLeft() *Node {
	fmt.Printf("Performing left rotation for node %v\n", root)
	if root == nil {
		return nil
	}

	// Make right node new root
	newRoot := root.right
	// Save left child of new root
	temp := newRoot.left
	// Make current root left child of new root
	newRoot.left = root
	// make temp right child of current root
	root.right = temp
	// Return new root
	return newRoot
}

// Rotate right
func (root *Node) RotateRight() *Node {
	fmt.Printf("Performing right rotation for node %v\n", root)
	if root == nil {
		return nil
	}

	// Make left node new root
	newRoot := root.left
	// Save right child node of new root
	temp := newRoot.right
	// Make current root right child node of new root
	newRoot.right = root
	// Make temp left child of current root
	root.left = temp
	// Return new root node
	return newRoot
}

// Get Balance factor for the given node
func (node *Node) GetBalanceFactor() int {
	if node == nil {
		return -1
	}

	return node.left.Height() - node.right.Height()
}

// Height of AVL tree
func (node *Node) Height() int {
	// Base condition
	if node == nil {
		return -1
	}

	heightOfLeftSubTree := node.left.Height()
	heightOfRightSubTree := node.right.Height()
	if heightOfLeftSubTree > heightOfRightSubTree {
		return heightOfLeftSubTree + 1
	} else {
		return heightOfRightSubTree + 1
	}
}

// Print AVL tree
func (node *Node) Print(s string) {
	// Base case
	if node == nil {
		return
	}

	s = fmt.Sprintf("%s %s", s, SPACE)
	node.right.Print(s)
	fmt.Printf("%s  %d \n", s, node.value)
	node.left.Print(s)
}

func main() {
	fmt.Println("Implementing AVL tree")
	avl := NewAVL([]int{10, 30, 20, 50, 25, 15, 70, 55, 5, 80})
	fmt.Printf("Created avl tree with root node %v\n", avl.root)
}
