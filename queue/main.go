package main

import (
	"fmt"
)

func push(arr []int, n int) []int {
	arr = append(arr, n)
	return arr
}

func pop(arr []int, n int) []int {
	arr = arr[1:n]
	return arr
}

func main() {
	arr := []int{7, 4, 9}

	arr = push(arr, 2)
	arr = push(arr, 6)
	fmt.Println(arr)

	arr = pop(arr, len(arr))
	fmt.Println(arr)
}
