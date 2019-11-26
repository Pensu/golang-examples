package main

import (
	"fmt"
)

func main() {
	arr := [6]int{43, 56, 5, 8, 91, 2}
	unsorted := len(arr) - 1
	sorted := false

	for sorted == false {
		sorted = true
		for i := 0; i <= unsorted-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				sorted = false
			}
		}
		unsorted--
	}

	fmt.Println(arr)
}
