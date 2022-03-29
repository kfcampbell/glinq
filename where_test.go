package main

import (
	"testing"
)

func TestWhere(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		pred     func(elem int) bool
		expected []int
	}{
		{
			name:  "evens",
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			pred: func(elem int) bool {
				return elem%2 == 0
			},
			expected: []int{0, 2, 4, 6, 8},
		},
	}

	for _, tc := range cases {
		actualWhere := Where(tc.input, tc.pred)
		if !sliceValueEquality(actualWhere, tc.expected) {
			t.Errorf("TestWhere %v: expected %v, got %v", tc.name, tc.expected, actualWhere)
		}

		ch := sliceToChan(tc.input)
		actualWhereCh := WhereCh(ch, tc.pred)
		out := chanToSlice(actualWhereCh)
		if !sliceValueEquality(out, tc.expected) {
			t.Errorf("TestWhere WhereCh %v: expected %v, got %v", tc.name, tc.expected, out)
		}
	}
}
