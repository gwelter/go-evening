package main

import (
	"errors"
	"fmt"
)

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot devide by 0")
	}
	return x / y, nil
}

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}
