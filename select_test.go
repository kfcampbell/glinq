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
