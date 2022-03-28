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

func TestFindInt(t *testing.T) {
	cases := []struct {
		name        string
		input       []int
		pred        func(elem int) bool
		expected    int
		expectedErr bool
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			func(elem int) bool {
				return elem == 2
			},
			2,
			false,
		},
		{
			"expandedCase",
			[]int{1, 4, 3, 23, 412, 19, 49, 67, 76, 15, 23, 18},
			func(elem int) bool {
				return elem == 412
			},
			412,
			false,
		},
		{
			"doesNotExist",
			[]int{1, 4, 3},
			func(elem int) bool {
				return elem == 412
			},
			412,
			true,
		},
		{
			"emptyList",
			[]int{},
			func(elem int) bool {
				return elem == 412
			},
			412,
			true,
		},
	}
	for _, tc := range cases {
		actual, err := Find(tc.input, tc.pred)
		if tc.expectedErr && err == nil {
			t.Errorf("TestFindInt %v: expected an error and got nil", tc.name)
		}

		if !tc.expectedErr && err != nil {
			t.Errorf("TestFindInt %v: did not expect error but got %v", tc.name, err)
		}

		if !tc.expectedErr && actual != tc.expected {
			t.Errorf("TestFindInt %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestFindString(t *testing.T) {
	cases := []struct {
		name        string
		input       []string
		pred        func(elem string) bool
		expected    string
		expectedErr bool
	}{
		{
			"happyCase",
			[]string{"abc", "def", "hijk"},
			func(elem string) bool {
				return elem == "def"
			},
			"def",
			false,
		},
		{
			"expandedCase",
			[]string{"abc", "def", "hijk", "lmnop", "qrs", "tuv", "wxy", "z"},
			func(elem string) bool {
				return elem == "hijk"
			},
			"hijk",
			false,
		},
		{
			"doesNotExist",
			[]string{"abc", "def", "hijk"},
			func(elem string) bool {
				return elem == "lmnop"
			},
			"lmnop",
			true,
		},
		{
			"emptyList",
			[]string{},
			func(elem string) bool {
				return elem == "def"
			},
			"def",
			true,
		},
	}
	for _, tc := range cases {
		actual, err := Find(tc.input, tc.pred)
		if tc.expectedErr && err == nil {
			t.Errorf("TestFindString %v: expected an error and got nil", tc.name)
		}

		if !tc.expectedErr && err != nil {
			t.Errorf("TestFindString %v: did not expect error but got %v", tc.name, err)
		}

		if !tc.expectedErr && actual != tc.expected {
			t.Errorf("TestFindString %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}
