package main

import (
	"fmt"
)

func main() {
	input := []uint{2, 3, 5, 19, 22, 7, 23, 5}
	i := IndexOf(input, 19)
	k := IndexOf(input, 8)
	l := LastIndexOf(input, 5)

	fmt.Printf("Index of i: %v, index of k: %v, last index of l: %v", i, k, l)
}
