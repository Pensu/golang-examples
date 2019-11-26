package main

import (
	"fmt"
)

func quicksort(k int, arr []int, left int, right int) {
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

	if arr[k] < arr[left] {
		quicksort(k, arr, 0, left-1)
	} else if arr[k] > arr[left] {
		quicksort(k, arr, left+1, right)
	} else {
		fmt.Println(arr, arr[k])
	}
}

func main() {
	arr := []int{6, 3, 8, 2, 9, 1, 5}
	quicksort(1, arr, 0, len(arr)-1)
	fmt.Println(arr)
}
