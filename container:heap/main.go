package main

import (
	"container/heap"
	"fmt"
)

type heapInt []int

func (n *heapInt) Pop() interface{} {
	old := *n
	val := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return val
}

func (n *heapInt) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n heapInt) Len() int {
	return len(n)
}

func (n heapInt) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapInt) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapInt{4, 7, 2, 9, 1, 5}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(3)
	myHeap.Push(6)
	fmt.Printf("Heap size: %d", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)

}
