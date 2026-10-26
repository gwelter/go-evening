package main

import "fmt"

func Swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}
