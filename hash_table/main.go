package main

import "fmt"

// SIZE is size of the table
const SIZE = 15

// Node denotes a node structure
type Node struct {
	data int
	next *Node
}

// HashTable denotes a hash table
type HashTable struct {
	table map[int]*Node
	size  int
}

func hashfunction(i, size int) int {
	return i % size
}

func insert(hash *HashTable, value int) int {
	index := hashfunction(value, hash.size)
	element := Node{data: value, next: hash.table[index]}
	hash.table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.table {
		if hash.table[k] != nil {
			t := hash.table[k]
			for t != nil {
				fmt.Printf("%d ->", t.data)
				t = t.next
			}
			fmt.Println()
		}
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{table: table, size: SIZE}

	for i := 0; i < 120; i++ {
		insert(hash, i)
	}

	traverse(hash)
}
