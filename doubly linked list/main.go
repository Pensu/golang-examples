package main

import "fmt"

// Node denotes a node structure
type Node struct {
	data  int
	left  *Node
	right *Node
}

var root = new(Node)

func insert(t *Node, n int) int {
	if root == nil {
		t = &Node{n, nil, nil}
		root = t
		return 0
	}

	if t.data == n {
		fmt.Println("Data exists")
		return -1
	}

	if t.right == nil {
		t.right = &Node{n, t, nil}
		return -2
	}

	return insert(t.right, n)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("Empty list")
	}

	for t != nil {
		fmt.Printf("%d ->", t.data)
		t = t.right
	}

	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("Empty list")
	}

	temp := t
	for t != nil {
		temp = t
		t = t.right
	}

	for temp.left != nil {
		fmt.Printf("%d ->", temp.data)
		temp = temp.left
	}

	fmt.Printf("%d ->", temp.data)
	fmt.Println()
}

func main() {
	root = nil
	insert(root, 2)
	insert(root, 5)
	insert(root, 7)
	insert(root, 3)

	traverse(root)
	reverse(root)
}
