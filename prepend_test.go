package main

import "testing"

func TestPrepend(t *testing.T) {
	cases := []struct {
		name     string
		source   []int
		elem     int
		expected []int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			0,
			[]int{0, 1, 2, 3},
		},
		{
			"emptyCase",
			[]int{},
			0,
			[]int{0},
		},
		{
			"longCase",
			[]int{42, 11, 29, 19, 17, 18, 25, 43, 99},
			11,
			[]int{11, 42, 11, 29, 19, 17, 18, 25, 43, 99},
		},
	}

	for _, tc := range cases {
		actual := Prepend(tc.source, tc.elem)
		success := sliceValueEquality(actual, tc.expected)
		if !success {
			t.Errorf("TestPrepend %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.source)
		actualCh := PrependCh(ch, tc.elem)
		actualChSl := chanToSlice(actualCh)
		successCh := sliceValueEquality(actualChSl, tc.expected)
		if !successCh {
			t.Errorf("TestPrepend PrependCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}
