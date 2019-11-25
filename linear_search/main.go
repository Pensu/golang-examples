package main

import (
	"fmt"
)

func main() {
	arr := [6]int{2, 3, 4, 5, 6, 7}
	searchElem := 4

	for i := 0; i < len(arr); i++ {
		if arr[i] == searchElem {
			fmt.Println(i)
			break
		} else if arr[i] > searchElem {
			break
		}
	}
}
