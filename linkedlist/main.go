package main

import "fmt"

// Node denotes a node structure
type Node struct {
	data     int
	nextNode *Node
}

func add(node *Node, val int) {
	newNode := Node{val, nil}
	for node.nextNode != nil {
		node = node.nextNode
	}

	node.nextNode = &newNode
}

func printlist(node *Node) {
	for node.nextNode != nil {
		fmt.Printf("%d-", node.data)
		node = node.nextNode
	}

	fmt.Printf("%d\n", node.data)
}

func insert(node *Node, newNode Node, k int) {
	for i := 1; i < k; i++ {
		node = node.nextNode
	}

	newNode.nextNode = node.nextNode
	node.nextNode = &newNode
}

func delete(node *Node, k int) {
	for i := 1; i < k-1; i++ {
		node = node.nextNode
	}

	node.nextNode = node.nextNode.nextNode
}

func main() {
	node1 := Node{3, nil}
	arr := []int{5, 7, 6, 2, 8}
	for _, val := range arr {
		add(&node1, val)
	}
	node2 := Node{9, nil}
	printlist(&node1)
	insert(&node1, node2, 3)
	printlist(&node1)
	delete(&node1, 5)
	printlist(&node1)
}
