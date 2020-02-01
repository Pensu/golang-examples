package main

import (
	"fmt"
)

type Node struct {
	value int
	connect []*Node
}


func dfs (node *Node) {
	if node == nil {
		return
	} else {
		fmt.Println(node.value)
		for _, i := range node.connect {
			dfs(i)
		}
	}
}

func main() {
	n1 := Node{1, nil}
	n2 := Node{2, nil}
	n3 := Node{3, nil}
	n4 := Node{4, nil}
	n5 := Node{5, nil}
	n6 := Node{6, nil}
	n7 := Node{7, nil}
	n8 := Node{8, nil}
	n9 := Node{9, nil}
	n10 := Node{10, nil}
	
	n1.connect = append(n1.connect, &n2)
	n1.connect = append(n1.connect, &n3)
	
	n2.connect = append(n2.connect, &n4)
	
	n3.connect = append(n3.connect, &n5)
	
	n4.connect = append(n4.connect, &n6)
	
	n5.connect = append(n5.connect, &n7)
	
	n4.connect = append(n4.connect, &n8)
	
	n2.connect = append(n2.connect, &n9)
	
	n4.connect = append(n4.connect, &n10)
	
	
	dfs(&n1)
} 

