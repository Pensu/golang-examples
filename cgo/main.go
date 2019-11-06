package main

// Go can understand these comments.

//#include <stdio.h>
//void callC() {
//	printf("Calling C code");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("Printing Go Code")
	C.callC()
	fmt.Println("Printing another Go Code")
}
