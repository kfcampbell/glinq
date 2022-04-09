package main

import (
	"testing"
)

func TestIntersectInt(t *testing.T) {
	cases := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "happyCase",
			first:    []int{1, 2, 3, 4},
			second:   []int{2, 3},
			expected: []int{2, 3},
		},
		{
			name:     "repeatingCase",
			first:    []int{6, 4, 5, 5, 8, 6, 7, 9, 10, 8},
			second:   []int{3, 2, 1, 5, 5, 8, 10},
			expected: []int{5, 8, 10},
		},
		{
			name:     "emptyFirst",
			first:    []int{},
			second:   []int{3, 2, 1},
			expected: []int{},
		},
		{
			name:     "emptySecond",
			first:    []int{3, 2, 2, 1},
			second:   []int{},
			expected: []int{},
		},
	}

	for _, tc := range cases {
		actual := Intersect(tc.first, tc.second)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestIntersect %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		firstCh := sliceToChan(tc.first)
		secondCh := sliceToChan(tc.second)
		actualCh := IntersectCh(firstCh, secondCh)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestIntersect IntersectCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}

func TestIntersectByInt(t *testing.T) {
	cases := []struct {
		name     string
		first    []int
		second   []int
		key      func(elem int) int
		expected []int
	}{
		{
			name:   "happyCase",
			first:  []int{1, 2, 3, 4, 5},
			second: []int{2, 3},
			key: func(elem int) int {
				return elem - 2
			},
			expected: []int{4, 5},
		},
		{
			name:   "repeatingCase",
			first:  []int{6, 4, 5, 5, 8, 6, 7, 9, 10, 8}, // 3, 2, 2, 2, 4, 3, 3, 4, 5, 4
			second: []int{3, 2, 1, 5, 5, 8, 10},
			key: func(elem int) int {
				return elem / 2
			},
			expected: []int{6, 4, 10},
		},
		{
			name:   "emptyFirst",
			first:  []int{},
			second: []int{3, 2, 1},
			key: func(elem int) int {
				return elem - 2
			},
			expected: []int{},
		},
		{
			name:   "emptySecond",
			first:  []int{3, 2, 2, 1},
			second: []int{},
			key: func(elem int) int {
				return elem - 2
			},
			expected: []int{},
		},
	}

	for _, tc := range cases {
		actual := IntersectBy(tc.first, tc.second, tc.key)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("IntersectBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		firstCh := sliceToChan(tc.first)
		secondCh := sliceToChan(tc.second)
		actualCh := IntersectByCh(firstCh, secondCh, tc.key)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("IntersectBy IntersectByCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}
