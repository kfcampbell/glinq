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

func TestSkip(t *testing.T) {
	cases := []struct {
		name     string
		source   []int
		skip     int
		expected []int
	}{
		{"happyCase", []int{0, 1, 2, 3, 4}, 2, []int{2, 3, 4}},
		{"empty", []int{}, 0, []int{}},
	}
	for _, tc := range cases {
		result := Skip(tc.source, tc.skip)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestSkip %v: expected %v, got %v", tc.name, tc.expected, result)
		}
		resultCh := SkipCh(sliceToChan(tc.source), tc.skip)
		resultChSlice := chanToSlice(resultCh)
		if !sliceValueEquality(resultChSlice, tc.expected) {
			t.Errorf("TestSkipCh %v: expected %v, got %v", tc.name, tc.expected, resultCh)
		}
	}
}

func TestSkipLast(t *testing.T){
	cases := []struct {
		name     string
		source   []int
		count    int
		expected []int
	}{
		{"happyCase", []int{0, 1, 2, 3, 4}, 2, []int{0, 1, 2}},
		{"empty", []int{}, 0, []int{}},
	}
	for _, tc := range cases {
		result := SkipLast(tc.source, tc.count)
		if !sliceValueEquality(result, tc.expected) {
			t.Errorf("TestSkipLast %v: expected %v, got %v", tc.name, tc.expected, result)
		}
		resultCh := SkipLastCh(sliceToChan(tc.source), tc.count)
		resultChSlice := chanToSlice(resultCh)
		if !sliceValueEquality(resultChSlice, tc.expected) {
			t.Errorf("TestSkipLastCh %v: expected %v, got %v", tc.name, tc.expected, resultCh)
		}
	}
}
