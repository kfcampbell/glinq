package main

import (
	"testing"
)

func sliceValueEquality[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func chanToSlice[T any](in <-chan T) []T {
	res := make([]T, 0)
	for {
		v, ok := <-in
		if !ok {
			return res
		}
		res = append(res, v)
	}
}

func TestSelect(t *testing.T) {
	type simple struct {
		id   int
		desc string
	}
	cases := []struct {
		name     string
		input    []simple
		selector func(obj simple) int
		expected []int
	}{
		{
			name: "happyCase",
			input: []simple{
				{
					id:   1,
					desc: "test one",
				},
				{
					id:   2,
					desc: "test two",
				},
			},
			selector: func(obj simple) int {
				return obj.id
			},
			expected: []int{1, 2},
		},
		{
			name: "expandedCase",
			input: []simple{
				{
					id:   0,
					desc: "test zero",
				},
				{
					id:   1,
					desc: "test one",
				},
				{
					id:   2,
					desc: "test two",
				},
				{
					id:   3,
					desc: "test three",
				},
				{
					id:   4,
					desc: "test four",
				},
				{
					id:   5,
					desc: "test five",
				},
			},
			selector: func(obj simple) int {
				return obj.id
			},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tc := range cases {
		actualSelected := Select(tc.input, tc.selector)
		if !sliceValueEquality(actualSelected, tc.expected) {
			t.Errorf("SelectTest error: expected slices to be equal. wanted %v, got %v", tc.expected, actualSelected)
		}

		ch := sliceToChan(tc.input)
		actualSelectedCh := SelectCh(ch, tc.selector)
		res := chanToSlice(actualSelectedCh)
		if !sliceValueEquality(res, tc.expected) {
			t.Errorf("SelectTest SelectCh error: expected slices to be equal. wanted %v, got %v", tc.expected, actualSelectedCh)
		}

	}
}
