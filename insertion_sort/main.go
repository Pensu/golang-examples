package main

import "fmt"

func main() {
	arr := [6]int{34, 32, 76, 4, 9, 24}

	for i := 1; i < len(arr); i++ {
		pos := i
		val := arr[i]

		for pos > 0 && arr[pos-1] > val {
			arr[pos] = arr[pos-1]
			pos--
		}
		arr[pos] = val
	}

	fmt.Println(arr)
}
