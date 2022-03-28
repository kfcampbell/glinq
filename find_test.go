package main

import (
	"testing"
)

func TestIndexOfInt(t *testing.T) {
	cases := []struct {
		info     string
		input    []int
		elem     int
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			2,
			1,
		},
		{
			"emptyCase",
			[]int{},
			5,
			-1,
		},
		{
			"biggerCase",
			[]int{1, 8, 9, 23, 44, 29, 97, 22, 42, 89, 76},
			89,
			9,
		},
		{
			"firstOfRepeating",
			[]int{0, 9, 9, 13, 2, 4},
			9,
			1,
		},
	}

	for _, tc := range cases {
		actual := IndexOf(tc.input, tc.elem)
		if actual != tc.expected {
			t.Errorf("%v: expected %v, got %v", tc.info, tc.elem, actual)
		}
	}
}
