package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	str := os.Args[1]

	t, _ := time.Parse("15:04", str)

	fmt.Println(t.Hour(), t.Minute())
}
