package main

import (
	"testing"
)

func TestIndexOfInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		elem     int
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			2,
			1,
		},
		{
			"emptyCase",
			[]int{},
			5,
			-1,
		},
		{
			"biggerCase",
			[]int{1, 8, 9, 23, 44, 29, 97, 22, 42, 89, 76},
			89,
			9,
		},
		{
			"firstOfRepeating",
			[]int{0, 9, 9, 13, 2, 4},
			9,
			1,
		},
		{
			"doesNotExist",
			[]int{8, 9, 10, 11, 12},
			14,
			-1,
		},
	}

	for _, tc := range cases {
		actual := IndexOf(tc.input, tc.elem)
		if actual != tc.expected {
			t.Errorf("TestIndexOfInt %v: expected %v, got %v", tc.name, tc.elem, actual)
		}
	}
}

func TestIndexOfString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		elem     string
		expected int
	}{
		{
			"happyCase",
			[]string{"abc", "def", "hijk"},
			"def",
			1,
		},
		{
			"doesNotExist",
			[]string{"abc", "def", "hijk"},
			"lmnop",
			-1,
		},
	}

	for _, tc := range cases {
		actual := IndexOf(tc.input, tc.elem)
		if actual != tc.expected {
			t.Errorf("TestIndexOfString %v: expected %v, got %v", tc.name, tc.elem, actual)
		}
	}
}

func TestLastIndexOfInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		elem     int
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			2,
			1,
		},
		{
			"twice",
			[]int{1, 2, 3, 4, 5, 4},
			4,
			5,
		},
		{
			"thrice",
			[]int{1, 2, 2, 2, 4},
			2,
			3,
		},
		{
			"doesNotExist",
			[]int{4, 18, 19, 22},
			42,
			-1,
		},
	}

	for _, tc := range cases {
		actual := LastIndexOf(tc.input, tc.elem)
		if actual != tc.expected {
			t.Errorf("TestLastIndexOfInt %v: expected %v, got %v", tc.name, tc.elem, actual)
		}
	}
}

func TestLastIndexOfString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		elem     string
		expected int
	}{
		{
			"happyCase",
			[]string{"abc", "def", "hijk"},
			"def",
			1,
		},
		{
			"twice",
			[]string{"abc", "def", "def", "hijk"},
			"def",
			2,
		},
		{
			"thrice",
			[]string{"abc", "def", "def", "def", "hijk"},
			"def",
			3,
		},
		{
			"doesNotExist",
			[]string{"abc", "def", "hijk"},
			"lmnop",
			-1,
		},
	}

	for _, tc := range cases {
		actual := LastIndexOf(tc.input, tc.elem)
		if actual != tc.expected {
			t.Errorf("TestLastIndexOfString %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}
