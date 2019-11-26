package main

import (
	"fmt"
)

// Node denotes a node structure
type Node struct {
	data     int
	nextNode *Node
}

func read(node Node, n int) {
	for i := 0; i <= n; i++ {
		node = *node.nextNode
	}

	fmt.Println(node.data)
}

func search(node Node, n int) {
	i := 1
	for node.nextNode != nil {
		if node.data == n {
			fmt.Println(i)
			break
		} else {
			i++
			node = *node.nextNode
		}
	}
}

func insert(node Node, newNode Node, k int) {
	for i := 0; i < k; i++ {
		node = *node.nextNode
	}

	nodeNext := node.nextNode

	newNode.nextNode = nodeNext
	node.nextNode = &newNode
}

func delete(node Node, n int) {
	for i := 0; i < n; i++ {
		node = *node.nextNode
	}

	nodeNext := *node.nextNode

	node.nextNode = nodeNext.nextNode
	nodeNext.nextNode = nil
}

func main() {
	node1 := Node{2, nil}
	node2 := Node{5, nil}
	node3 := Node{6, nil}
	node4 := Node{3, nil}

	node1.nextNode = &node2
	node2.nextNode = &node3
	node3.nextNode = &node4
	fmt.Println(node1, node2, node3)

	read(node1, 1)
	search(node1, 6)

	node5 := Node{7, nil}
	insert(node1, node5, 2)
	delete(node1, 2)
}
