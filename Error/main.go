package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Sending you an error!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("No errors found")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("No errors found")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Sending you an error!" {
		fmt.Println("!!")
	}
}
