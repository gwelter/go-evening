package main

import "fmt"

var allocated = 0

func AllocateBuffer() *string {
	if allocated >= 3 {
		return nil
	}
	allocated++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
