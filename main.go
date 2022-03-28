package main

import (
	"fmt"
)

func main() {
	list := []uint{2, 3, 5, 19, 22, 7, 23, 5}
	i := IndexOf(list, 19)
	k := IndexOf(list, 8)
	l := LastIndexOf(list, 5)

	fmt.Printf("Index of i: %v, index of k: %v, last index of l: %v", i, k, l)
}
