package main

import (
	"fmt"
)

func quicksort(arr []int, left int, right int) {
	if right-left <= 0 {
		return
	}

	pivot := right
	for true {
		for arr[left] < arr[pivot] {
			left++
		}

		for arr[right] > arr[pivot] {
			right--
		}

		if right <= left {
			break
		}
		temp := arr[left]
		arr[left] = arr[pivot]
		arr[pivot] = temp
	}

	quicksort(arr, 0, left-1)

	quicksort(arr, left+1, right)
}

func main() {
	arr := []int{6, 3, 8, 2, 9, 1, 5}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
