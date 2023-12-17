package main

import "fmt"

const SPACE string = "     "

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
	fmt.Printf("Removing %d from BST\n", v)
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
		fmt.Printf("Found %d at node %v\n", v, node)
		// When node has 0 child
		if node.left == nil && node.right == nil {
			fmt.Printf("0 child: Removing node %v and returning %v\n", node, nil)
			node = nil
			return nil
		}

		// When node has 1 left child
		if node.left != nil && node.right == nil {
			fmt.Printf("1 left child: Removing node %v and returning %v\n", node, node.left)
			temp := node.left
			node = nil
			return temp
		}

		// When node has 1 right child
		if node.left == nil && node.right != nil {
			fmt.Printf("1 right child: Removing node %v and returning %v\n", node, node.right)
			temp := node.right
			node = nil
			return temp
		}

		// When node has 2 children
		if node.left != nil && node.right != nil {
			minNode := node.right.Min()
			node.value = minNode.value
			node.right = node.right.Remove(minNode.value)
			fmt.Printf("2 children: Removing node %v and returning %v\n", minNode, node)
			return node
		}
	} else if v < node.value {
		// Node to be removed is in left part of the tree
		fmt.Printf("Removing %d from left part of tree with root node %v\n", v, node.left)
		node.left = node.left.Remove(v)
	} else {
		// Node to be removed is in right part of the tree
		fmt.Printf("Removing %d from right part of tree with root node %v\n", v, node.right)
		node.right = node.right.Remove(v)
	}

	return node
}

// Search given value in the BST(recursive approach)
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

// Search given value in the BST(iterative approach)
func (node *Node) SearchIterative(v int) *Node {
	temp := node
	for temp != nil {
		if v == temp.value {
			return temp
		} else if v < temp.value {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	return temp
}

// Print BST
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

// Height of BST
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

// In-order traverser of BST - DFS (Depth First Search) technique
func (node *Node) InOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	node.left.InOrderTraverser()
	fmt.Printf("%d  ", node.value)
	node.right.InOrderTraverser()
}

// Pre-order traverser of the tree - DFS (Depth First Search) technique
func (node *Node) PreOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	fmt.Printf("%d  ", node.value)
	node.left.PreOrderTraverser()
	node.right.PreOrderTraverser()
}

// Post-order traverser of the tree - DFS (Depth First Search) technique
func (node *Node) PostOrderTraverser() {
	// Base case
	if node == nil {
		return
	}

	node.left.PostOrderTraverser()
	node.right.PostOrderTraverser()
	fmt.Printf("%d  ", node.value)
}

// Level oder traverser of the tree - BFS (Breadth First Search) technique
func (node *Node) LevelOrderTraverser() {
	fmt.Println("Level order traverser of the tree")
	// Get the height of the tree
	height := node.Height()

	// Traverse each level iteratively
	for i := 0; i <= height; i++ {
		traverseTheLevel(node, i)
		fmt.Println()
	}
}

// Internal function of LevelOrderTraverser method
func traverseTheLevel(node *Node, level int) {
	// Base case 1
	if node == nil {
		return
	}

	// Base case 2
	if level == 0 {
		fmt.Printf("%d	", node.value)
		return
	}

	// Recursively traverse the left subtree for the given level
	traverseTheLevel(node.left, level-1)
	// Recursively traverse the right subtree for the given level
	traverseTheLevel(node.right, level-1)
}

func main() {
	fmt.Println("Implementing binary search tree")
	bst := NewBST([]int{9, 4, 6, 20, 170, 15, 1, 22, 200})
	fmt.Println("Printing the BST")
	bst.root.Print("")
	fmt.Printf("Height of BST is %d\n", bst.root.Height())
	bst.root.InOrderTraverser()
	fmt.Println()
	fmt.Printf("Recursive search 6 - %v\n", bst.root.Search(6))
	fmt.Printf("Iterative search 6 - %v\n", bst.root.SearchIterative(6))
	fmt.Printf("Max - %v\n", bst.root.Max())
	fmt.Printf("Min - %v\n", bst.root.Min())
	bst.Remove(9)
	bst.root.InOrderTraverser()
	fmt.Println()
	fmt.Printf("Recursive search 9 - %v\n", bst.root.Search(9))
	fmt.Printf("Iterative search 9 - %v\n", bst.root.SearchIterative(9))
	fmt.Printf("Max - %v\n", bst.root.Max())
	fmt.Printf("Min - %v\n", bst.root.Min())
	bst.root.PreOrderTraverser()
	fmt.Println()
	bst.root.PostOrderTraverser()
	fmt.Println()
	bst.root.LevelOrderTraverser()
	fmt.Println()
	fmt.Println("Printing the BST")
	bst.root.Print("")
}
