package main

import (
	"fmt"
)

type Node struct {
	value int
	left *Node
	right *Node
}

func inorder (node *Node) {
	if node == nil {
		return
	}
	
	inorder(node.left)
	fmt.Println(node.value)
	inorder(node.right)
}

func postorder (node *Node) {
	if node == nil {
		return
	}
	
	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.value)
}

func preorder (node *Node) {
	if node == nil {
		return
	}
	
	fmt.Println(node.value)
	preorder(node.left)
	preorder(node.right)
}

func main() {
	n1 := Node{1, nil, nil}
	n2 := Node{2, nil, nil}
	n3 := Node{3, nil, nil}
	n4 := Node{4, nil, nil}
	n5 := Node{5, nil, nil}
	n6 := Node{6, nil, nil}
	n7 := Node{7, nil, nil}
	n8 := Node{8, nil, nil}
	n9 := Node{9, nil, nil}
	n10 := Node{10, nil, nil}
	
	n1.left = &n2
	n1.right = &n3
	
	n2.left = &n4
	n2.right = &n5
	
	n3.left = &n6
	n3.right = &n7
	
	n4.left = &n8
	
	n5.right = &n9
	
	n6.left = &n10
	
	fmt.Println("Printing inorder")
	inorder(&n1)
	fmt.Println("Printing postorder")
	postorder(&n1)
	fmt.Println("Printing preorder")
	preorder(&n1)
} 
