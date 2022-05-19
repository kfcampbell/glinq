package main

import "testing"

func TestRange(t *testing.T) {
	cases := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{"happyCase", 0, 5, []int{0, 1, 2, 3, 4}},
		{"empty", 0, 0, []int{}},
	}
	for _, tc := range cases {
		result := Range(tc.start, tc.end)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestRange %v: expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestRepeat(t *testing.T) {
	cases := []struct {
		name     string
		elem     int
		repeat   int
		expected []int
	}{
		{"happyCase", 1, 5, []int{1, 1, 1, 1, 1}},
		{"empty", 0, 0, []int{}},
	}
	for _, tc := range cases {
		result := Repeat(tc.elem, tc.repeat)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestRepeat %v: expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}
