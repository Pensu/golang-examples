package main

import (
	"fmt"
)

func push(arr []int, n int) []int {
	arr = append(arr, n)
	return arr
}

func pop(arr []int, n int) []int {
	arr = arr[:n-1]
	return arr
}

func main() {
	arr := []int{4, 6, 2}

	arr = push(arr, 5)
	arr = push(arr, 7)
	fmt.Println(arr)

	arr = pop(arr, len(arr))
	fmt.Println(arr)
}
