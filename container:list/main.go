package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()

	values.PushBack(3)
	values.PushBack(2)
	values.PushBack(7)
	values.PushBack(5)

	printList(values)

}
