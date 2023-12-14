package main

import "fmt"

// BST structure
type BST struct {
	root *Node
}

// BST node
type Node struct {
	value int
	left  *Node
	right *Node
}

// New bst
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

// Insert into BST
func (b *BST) Insert(v int) {
	b.root = insertIntoBST(b.root, v)
}

func insertIntoBST(root *Node, v int) *Node {
	// Base condition
	if root == nil {
		node := NewNode(v)
		return node
	}

	if v > root.value {
		root.right = insertIntoBST(root.right, v)
	} else {
		root.left = insertIntoBST(root.left, v)
	}

	return root
}

// Search from BST
func (b *BST) Search(root *Node, v int) *Node {
	// Base condition 1
	if root == nil {
		return nil
	}

	// Base condition 2
	if v == root.value {
		return root
	}

	if v > root.value {
		return b.Search(root.right, v)
	} else {
		return b.Search(root.left, v)
	}
}

// Remove from BST
func (b *BST) Remove(root *Node, v int) *Node {
	if root == nil {
		return nil
	}

	// Found the node to be removed
	if root.value == v {
		// When node has 0 child
		if root.left == nil && root.right == nil {
			root = nil
			return nil
		}

		// When node has 1 left child
		if root.left != nil && root.right == nil {
			temp := root.left
			root = nil
			return temp
		}

		// When node has 1 right child
		if root.left == nil && root.right != nil {
			temp := root.right
			root = nil
			return temp
		}

		// When node has 2 children
		if root.left != nil && root.right != nil {
			minNode := b.Min(root.right)
			root.value = minNode.value
			root.right = b.Remove(root, minNode.value)
			return root
		}
	} else if v < root.value {
		// Node to be removed is in left part of the tree
		root.left = b.Remove(root.left, v)
	} else {
		// Node to be removed is in right part of the tree
		root.right = b.Remove(root.right, v)
	}

	return root
}

// In-order traverser of BST
func (b *BST) InOrderTraverser(node *Node) {
	// Base case
	if node == nil {
		return
	}

	b.InOrderTraverser(node.left)
	fmt.Printf("%d  ", node.value)
	b.InOrderTraverser(node.right)
}

// Min from BST
func (b *BST) Min(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}

	return node
}

// Max from BST
func (b *BST) Max(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}

	return node
}

func main() {
	fmt.Println("Implementing binary search tree")
	bst := NewBST([]int{9, 4, 6, 20, 170, 15, 1})
	bst.InOrderTraverser(bst.root)
	fmt.Println()
	fmt.Printf("Search 6 - %v\n", bst.Search(bst.root, 6))
	fmt.Printf("Max - %v\n", bst.Max(bst.root))
	fmt.Printf("Min - %v\n", bst.Min(bst.root))
	bst.Remove(bst.root, 200)
	bst.InOrderTraverser(bst.root)
	fmt.Println()
	fmt.Printf("Search 6 - %v\n", bst.Search(bst.root, 6))
	fmt.Printf("Max - %v\n", bst.Max(bst.root))
	fmt.Printf("Min - %v\n", bst.Min(bst.root))
}