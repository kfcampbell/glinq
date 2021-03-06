package main

import (
	"testing"
)

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

func TestChunk(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		size     int
		expected [][]int
	}{
		{
			"happyEvenCase",
			[]int{1, 2, 3, 4, 5, 6},
			2,
			[][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
		},
		{
			"happyOddCase",
			[]int{1, 2, 3, 4, 5},
			2,
			[][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
		{
			"emptyCase",
			[]int{},
			8,
			[][]int{},
		},
		{
			"oneChunk",
			[]int{1, 2, 3},
			3,
			[][]int{
				{1, 2, 3},
			},
		},
	}

	for _, tc := range cases {
		actualChunk := Chunk(tc.input, tc.size)
		if !sliceOfSliceValueEquality(actualChunk, tc.expected) {
			t.Errorf("TestChunk %v: expected %v, got %v", tc.name, tc.expected, actualChunk)
		}

		ch := sliceToChan(tc.input)
		actualChunkCh := ChunkCh(ch, tc.size)
		result := chanToSlice(actualChunkCh)
		if !sliceOfSliceValueEquality(result, tc.expected) {
			t.Errorf("TestChunk ChunkCh %v: expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestDistinct(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"happyCase",
			[]int{2, 3, 4, 2, 2, 3, 5},
			[]int{2, 3, 4, 5},
		},
		{
			"emptyCase",
			[]int{},
			[]int{},
		},
	}

	for _, tc := range cases {
		actual := Distinct(tc.input)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestDistinct %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.input)
		actualCh := DistinctCh(ch)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestDistinct DistinctCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}

func TestDistinctString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			"happyCase",
			[]string{"happy", "happy", "birthday"},
			[]string{"happy", "birthday"},
		},
		{
			"emptyCase",
			[]string{},
			[]string{},
		},
	}

	for _, tc := range cases {
		actual := Distinct(tc.input)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestDistinct %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.input)
		actualCh := DistinctCh(ch)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestDistinct DistinctCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}

func TestDistinctBy(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
		key      func(elem int) int
	}{
		{
			"happyCase",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2},
			func(elem int) int {
				return elem % 2
			},
		},
		{
			"emptyCase",
			[]int{},
			[]int{},
			func(elem int) int {
				return elem % 2
			},
		},
	}

	for _, tc := range cases {
		actual := DistinctBy(tc.input, tc.key)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestDistinctBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.input)
		actualCh := DistinctByCh(ch, tc.key)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestDistinctBy DistinctByCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}

}

func TestExceptFloat(t *testing.T) {
	cases := []struct {
		name     string
		first    []float64
		second   []float64
		expected []float64
	}{
		{
			name:     "happyCase",
			first:    []float64{2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5},
			second:   []float64{2.2},
			expected: []float64{2.0, 2.1, 2.3, 2.4, 2.5},
		},
		{
			name:     "emptyFirstCase",
			first:    []float64{},
			second:   []float64{2.2},
			expected: []float64{},
		},
		{
			name:     "emptySecondCase",
			first:    []float64{2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5},
			second:   []float64{},
			expected: []float64{2.0, 2.1, 2.2, 2.3, 2.4, 2.5},
		},
	}
	for _, tc := range cases {
		actual := Except(tc.first, tc.second)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestExcept %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		first := sliceToChan(tc.first)
		second := sliceToChan(tc.second)
		actualCh := ExceptCh(first, second)
		actualChSl := chanToSlice(actualCh)

		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestExcept ExceptCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}

func TestExceptByInt(t *testing.T) {
	cases := []struct {
		name     string
		first    []int
		second   []int
		key      func(int) int
		expected []int
	}{
		{
			name:   "simpleCase",
			first:  []int{1, 2, 3},
			second: []int{2},
			key: func(i int) int {
				return i % 2
			},
			expected: []int{1},
		},
		{
			name:   "expandedCase",
			first:  []int{1, 2, 2, 3, 4, 5, 5, 6, 7},
			second: []int{2, 4, 6},
			key: func(i int) int {
				return i % 2
			},
			expected: []int{1},
		},
		{
			name:   "emptyFirstCase",
			first:  []int{},
			second: []int{2, 4, 6},
			key: func(i int) int {
				return i % 2
			},
			expected: []int{},
		},
		{
			name:   "emptySecondCase",
			first:  []int{1, 2, 3, 4, 5, 6},
			second: []int{},
			key: func(i int) int {
				return i % 2
			},
			expected: []int{1, 2},
		},
		{
			name:   "anotherCase",
			first:  []int{2, 2, 3, 4, 4, 4, 7, 8},
			second: []int{2, 4},
			key: func(i int) int {
				return i / 2
			},
			expected: []int{7, 8},
		},
	}

	for _, tc := range cases {
		actual := ExceptBy(tc.first, tc.second, tc.key)
		if !sliceValueEquality(actual, tc.expected) {
			t.Errorf("TestExceptBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		first := sliceToChan(tc.first)
		second := sliceToChan(tc.second)
		actualCh := ExceptByCh(first, second, tc.key)
		actualChSl := chanToSlice(actualCh)
		if !sliceValueEquality(actualChSl, tc.expected) {
			t.Errorf("TestExceptBy ExceptByCh %v: expected %v, got %v", tc.name, tc.expected, actualChSl)
		}
	}
}
