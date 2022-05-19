package main

import "testing"

func TestPrependInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			4,
			[]int{4, 1, 2, 3},
		},
		{
			"empty",
			[]int{},
			4,
			[]int{4},
		},
	}

	for _, tc := range cases {
		result := Prepend(tc.input, tc.value)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestPrependInt %v: expected %v, got %v", tc.name, tc.expected, tc.input)
		}

		resultCh := PrependCh(sliceToChan(tc.input), tc.value)
		result = chanToSlice(resultCh)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestPrependIntCh %v: expected %v, got %v", tc.name, tc.expected, tc.input)
		}
	}
}
