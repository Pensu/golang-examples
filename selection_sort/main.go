package main

import "fmt"

func main() {
	arr := [6]int{43, 23, 9, 6, 12, 55}

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := arr[minIndex]
			arr[minIndex] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println(arr)
}
