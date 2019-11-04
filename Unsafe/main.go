package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64
	p1 := &value
	p2 := (*int32)(unsafe.Pointer(p1))
	*p1 = 65363517238971293
	fmt.Println(*p1, *p2)
}
