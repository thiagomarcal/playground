package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %v", result)
	}

}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}
