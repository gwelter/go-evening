package main

import (
	"fmt"
	"os"
)

func CheckFile(filename string) bool {
	_, err := os.ReadFile(filename)
	return err == nil
}

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}
