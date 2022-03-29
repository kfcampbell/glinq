package main

import (
	"testing"
)

func sliceToChan[T any](list []T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

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
		actIndexOf := IndexOf(tc.input, tc.elem)
		if actIndexOf != tc.expected {
			t.Errorf("TestIndexOfInt %v: expected %v, got %v", tc.name, tc.elem, actIndexOf)
		}

		ch := sliceToChan(tc.input)
		actIndexOfCh := IndexOfCh(ch, tc.elem)
		if actIndexOfCh != tc.expected {
			t.Errorf("TestIndexOfInt IndexOfCh %v: expected %v, got %v", tc.name, tc.elem, actIndexOfCh)
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
		actIndexOf := IndexOf(tc.input, tc.elem)
		if actIndexOf != tc.expected {
			t.Errorf("TestIndexOfString %v: expected %v, got %v", tc.name, tc.elem, actIndexOf)
		}

		ch := sliceToChan(tc.input)
		actIndexOfCh := IndexOfCh(ch, tc.elem)
		if actIndexOfCh != tc.expected {
			t.Errorf("TestIndexOfString IndexOfCh %v: expected %v, got %v", tc.name, tc.elem, actIndexOfCh)
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
		actLastIndexOf := LastIndexOf(tc.input, tc.elem)
		if actLastIndexOf != tc.expected {
			t.Errorf("TestLastIndexOfInt %v: expected %v, got %v", tc.name, tc.elem, actLastIndexOf)
		}

		ch := sliceToChan(tc.input)
		actLastIndexOfCh := LastIndexOfCh(ch, tc.elem)
		if actLastIndexOfCh != tc.expected {
			t.Errorf("TestLastIndexOfInt LastIndexOfCh %v: expected %v, got %v", tc.name, tc.elem, actLastIndexOfCh)
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
		actLastIndexOf := LastIndexOf(tc.input, tc.elem)
		if actLastIndexOf != tc.expected {
			t.Errorf("TestLastIndexOfString %v: expected %v, got %v", tc.name, tc.expected, actLastIndexOf)
		}

		ch := sliceToChan(tc.input)
		actLastIndexOfCh := LastIndexOfCh(ch, tc.elem)
		if actLastIndexOfCh != tc.expected {
			t.Errorf("TestLastIndexOfString LastIndexOfCh %v: expected %v, got %v", tc.name, tc.expected, actLastIndexOfCh)
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
			"emptySlice",
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
			"emptySlice",
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

func TestMinInt(t *testing.T) {
	cases := []struct {
		name        string
		input       []int
		expected    int
		expectedErr bool
	}{
		{
			"happyCase",
			[]int{14, 8, 9, 12},
			8,
			false,
		},
		{
			"emptySlice",
			[]int{},
			8,
			true,
		},
	}
	for _, tc := range cases {
		actual, err := Min(tc.input)
		if tc.expectedErr && err == nil {
			t.Errorf("TestMinInt %v: expected an error and got nil", tc.name)
		}

		if !tc.expectedErr && err != nil {
			t.Errorf("TestMinInt %v: did not expect error but got %v", tc.name, err)
		}

		if !tc.expectedErr && actual != tc.expected {
			t.Errorf("TestMinInt %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestMinCh(t *testing.T) {
	cases := []struct {
		name     string
		expected int
		fill     func(ch chan<- int)
	}{
		{
			name:     "happyCase",
			expected: 0,
			fill: func(ch chan<- int) {
				i := 0
				for i < 10 {
					ch <- i
					i++
				}
				close(ch)
			},
		},
		{
			name:     "nonTrivial",
			expected: -15,
			fill: func(ch chan<- int) {
				i := 0
				ch <- -10
				for i < 100 {
					ch <- (i * 2) % 3
					i++
				}
				ch <- -15
				close(ch)
			},
		},
	}

	for _, tc := range cases {
		ch := make(chan int)
		go tc.fill(ch)
		actual, err := MinCh(ch)
		if err != nil {
			t.Errorf("TestMinCh %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestMinChError(t *testing.T) {
	ch := make(chan int)
	close(ch)

	min, err := MinCh(ch)
	if err == nil {
		t.Errorf("TestMinChError: expected err, got nil and %v min", min)
	}
}
