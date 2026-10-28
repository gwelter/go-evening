package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func Keys(entries map[int]string) []int {
	keys := []int{}
	for key := range entries {
		keys = append(keys, key)
	}
	return keys
}

func Values(entries map[int]string) []string {
	values := []string{}
	for _, value := range entries {
		values = append(values, value)
	}
	return values
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}
