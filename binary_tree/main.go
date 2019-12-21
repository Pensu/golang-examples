package main

import "fmt"

// Tree struct denotes a tree node
type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func (t *Tree) insert(n int) *Tree {
	if t == nil {
		return &Tree{n, nil, nil}
	}

	if t.data == n {
		return t
	}

	if n < t.data {
		t.left = t.left.insert(n)
		return t
	}

	t.right = t.right.insert(n)
	return t
}

func (t *Tree) traverse() {
	if t == nil {
		return
	}

	t.left.traverse()
	fmt.Println(t.data)
	t.right.traverse()
}

func main() {
	t := &Tree{5, nil, nil}
	t.insert(2)
	t.insert(7)
	t.insert(4)
	t.insert(9)

	t.traverse()
}
