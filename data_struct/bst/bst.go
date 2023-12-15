package main

import "fmt"

// BST tree structure
type BST struct {
	root *Node
}

// BST node structure
type Node struct {
	value int
	left  *Node
	right *Node
}

// New BST tree
func NewBST(arr []int) *BST {
	bst := &BST{}
	for _, v := range arr {
		bst.Insert(v)
	}
	return bst
}

// New BST node
func NewNode(v int) *Node {
	return &Node{value: v}
}

// Wrapper insert into BST
func (b *BST) Insert(v int) {
	b.root = b.root.Insert(v)
}

// Wrapper remove from BST
func (b *BST) Remove(v int) {
	b.root = b.root.Remove(v)
}

// Internal insert into BST
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

	return node
}

// Internal remove from BST
func (node *Node) Remove(v int) *Node {
	if node == nil {
		return nil
	}

	// Found the node to be removed
	if node.value == v {
		// When node has 0 child
		if node.left == nil && node.right == nil {
			node = nil
			return nil
		}

		// When node has 1 left child
		if node.left != nil && node.right == nil {
			temp := node.left
			node = nil
			return temp
		}

		// When node has 1 right child
		if node.left == nil && node.right != nil {
			temp := node.right
			node = nil
			return temp
		}

		// When node has 2 children
		if node.left != nil && node.right != nil {
			minNode := node.right.Min()
			node.value = minNode.value
			node.right = node.Remove(minNode.value)
			return node
		}
	} else if v < node.value {
		// Node to be removed is in left part of the tree
		node.left = node.left.Remove(v)
	} else {
		// Node to be removed is in right part of the tree
		node.right = node.right.Remove(v)
	}

	return node
}

// Search from BST
func (node *Node) Search(v int) *Node {
	// Base condition 1
	if node == nil {
		return nil
	}

	// Base condition 2
	if v == node.value {
		return node
	}

	if v > node.value {
		return node.right.Search(v)
	} else {
		return node.left.Search(v)
	}
}

// In-order traverser of BST
func (node *Node) InOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	node.left.InOrderTraverser()
	fmt.Printf("%d  ", node.value)
	node.right.InOrderTraverser()
}

// Pre-order traverser of the tree
func (node *Node) PreOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	fmt.Printf("%d  ", node.value)
	node.left.PreOrderTraverser()
	node.right.PreOrderTraverser()
}

// Post-order traverser of the tree
func (node *Node) PostOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	node.left.PostOrderTraverser()
	node.right.PostOrderTraverser()
	fmt.Printf("%d  ", node.value)
}

// Min from BST
func (node *Node) Min() *Node {
	for node.left != nil {
		node = node.left
	}

	return node
}

// Max from BST
func (node *Node) Max() *Node {
	for node.right != nil {
		node = node.right
	}

	return node
}

func main() {
	fmt.Println("Implementing binary search tree")
	bst := NewBST([]int{9, 4, 6, 20, 170, 15, 1})
	bst.root.InOrderTraverser()
	fmt.Println()
	fmt.Printf("Search 6 - %v\n", bst.root.Search(6))
	fmt.Printf("Max - %v\n", bst.root.Max())
	fmt.Printf("Min - %v\n", bst.root.Min())
	bst.Remove(6)
	bst.root.InOrderTraverser()
	fmt.Println()
	fmt.Printf("Search 6 - %v\n", bst.root.Search(6))
	fmt.Printf("Max - %v\n", bst.root.Max())
	fmt.Printf("Min - %v\n", bst.root.Min())
	bst.root.PreOrderTraverser()
	fmt.Println()
	bst.root.PostOrderTraverser()
}
