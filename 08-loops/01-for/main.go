package main

import "fmt"

func Alphabet(length int) []string {
	var letters = []string{}
	for x := 0; x < length; x++ {
		letters = append(letters, characterByIndex(x))
	}
	return letters
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
