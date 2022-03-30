package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		pred     func(value int) bool
		expected bool
	}{
		{
			"evens",
			[]int{2, 4, 6, 8},
			func(value int) bool {
				return value%2 == 0
			},
			true,
		},
		{
			"evensNotExpected",
			[]int{1, 2, 3, 4, 5, 6},
			func(value int) bool {
				return value%2 == 0
			},
			false,
		},
	}

	for _, tc := range cases {
		actualAll := All(tc.input, tc.pred)
		if actualAll != tc.expected {
			t.Errorf("TestAll %v: expected %v, got %v", tc.name, tc.expected, actualAll)
		}

		ch := sliceToChan(tc.input)
		actualAllCh := AllCh(ch, tc.pred)
		if actualAllCh != tc.expected {
			t.Errorf("TestAll AllCh %v: expected %v, got %v", tc.name, tc.expected, actualAllCh)
		}
	}
}

func TestAny(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		pred     func(value int) bool
		expected bool
	}{
		{
			"evens",
			[]int{3, 5, 6, 9},
			func(value int) bool {
				return value%2 == 0
			},
			true,
		},
		{
			"evensNotExpected",
			[]int{1, 3, 5, 7},
			func(value int) bool {
				return value%2 == 0
			},
			false,
		},
	}

	for _, tc := range cases {
		actualAny := Any(tc.input, tc.pred)
		if actualAny != tc.expected {
			t.Errorf("TestAny %v: expected %v, got %v", tc.name, tc.expected, actualAny)
		}

		ch := sliceToChan(tc.input)
		actualAnyCh := AnyCh(ch, tc.pred)
		if actualAnyCh != tc.expected {
			t.Errorf("TestAny AnyCh %v: expected %v, got %v", tc.name, tc.expected, actualAnyCh)
		}
	}
}
