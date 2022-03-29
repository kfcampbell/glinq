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
		name     string
		input    []int
		pred     func(elem int) bool
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			func(elem int) bool {
				return elem == 2
			},
			2,
		},
		{
			"expandedCase",
			[]int{1, 4, 3, 23, 412, 19, 49, 67, 76, 15, 23, 18},
			func(elem int) bool {
				return elem == 412
			},
			412,
		},
	}
	for _, tc := range cases {
		actFind, err := Find(tc.input, tc.pred)

		if err != nil || actFind != tc.expected {
			t.Errorf("TestFindInt %v: expected %v, got %v, err %v", tc.name, tc.expected, actFind, err)
		}

		ch := sliceToChan(tc.input)
		actFindCh, err := FindCh(ch, tc.pred)
		if err != nil || actFindCh != tc.expected {
			t.Errorf("TestFindInt FindCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actFindCh, err)
		}
	}
}

func TestFindIntErr(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		pred  func(elem int) bool
	}{
		{
			"doesNotExist",
			[]int{1, 4, 3},
			func(elem int) bool {
				return elem == 412
			},
		},
		{
			"emptySlice",
			[]int{},
			func(elem int) bool {
				return elem == 412
			},
		},
	}

	for _, tc := range cases {
		actFind, err := Find(tc.input, tc.pred)
		if err == nil {
			t.Errorf("TestFindIntErr %v: wanted err but got nil, found :%v", tc.name, actFind)
		}
		ch := sliceToChan(tc.input)
		actFindCh, err := FindCh(ch, tc.pred)
		if err == nil {
			t.Errorf("TestFindIntErr FindCh %v: wanted err but got nil, found :%v", tc.name, actFindCh)
		}
	}
}

func TestFindString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		pred     func(elem string) bool
		expected string
	}{
		{
			"happyCase",
			[]string{"abc", "def", "hijk"},
			func(elem string) bool {
				return elem == "def"
			},
			"def",
		},
		{
			"expandedCase",
			[]string{"abc", "def", "hijk", "lmnop", "qrs", "tuv", "wxy", "z"},
			func(elem string) bool {
				return elem == "hijk"
			},
			"hijk",
		},
	}
	for _, tc := range cases {
		actFind, err := Find(tc.input, tc.pred)
		if err != nil || actFind != tc.expected {
			t.Errorf("TestFindString %v: expected %v, got %v, err %v", tc.name, tc.expected, actFind, err)
		}
		ch := sliceToChan(tc.input)
		actFindCh, err := FindCh(ch, tc.pred)
		if err != nil || actFindCh != tc.expected {
			t.Errorf("TestFindString FindCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actFindCh, err)
		}
	}
}

func TestMinInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"happyCase",
			[]int{14, 8, 9, 12},
			8,
		},
		{
			"negatives",
			[]int{-13, -11, -98, 45, 0, 199, -2},
			-98,
		},
	}
	for _, tc := range cases {
		actMin, err := Min(tc.input)

		if actMin != tc.expected || err != nil {
			t.Errorf("TestMinInt %v: expected %v, got %v, err: %v", tc.name, tc.expected, actMin, err)
		}

		ch := sliceToChan(tc.input)
		actMinCh, err := MinCh(ch)
		if actMinCh != tc.expected || err != nil {
			t.Errorf("TestMinInt MinCh %v: expected %v, got %v, err: %v", tc.name, tc.expected, actMinCh, err)
		}
	}
}

func TestMinError(t *testing.T) {
	list := make([]int, 0)

	min, err := Min(list)
	if err == nil {
		t.Errorf("TestMinError: expected err, got nil and %v min", min)
	}

	ch := sliceToChan(list)
	minCh, err := MinCh(ch)
	if err == nil {
		t.Errorf("TestMinError MinCh: expected err, got nil and %v min", minCh)
	}
}

func TestMaxInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"happyCase",
			[]int{14, 8, 9, 12},
			14,
		},
		{
			"negatives",
			[]int{-13, -11, -98, 45, 0, 199, -2},
			199,
		},
	}
	for _, tc := range cases {
		actMax, err := Max(tc.input)

		if actMax != tc.expected || err != nil {
			t.Errorf("TestMaxInt %v: expected %v, got %v, err: %v", tc.name, tc.expected, actMax, err)
		}

		ch := sliceToChan(tc.input)
		actMaxCh, err := MaxCh(ch)
		if actMaxCh != tc.expected || err != nil {
			t.Errorf("TestMaxInt MaxCh %v: expected %v, got %v, err: %v", tc.name, tc.expected, actMaxCh, err)
		}
	}
}

func TestMaxError(t *testing.T) {
	list := make([]int, 0)

	max, err := Max(list)
	if err == nil {
		t.Errorf("TestMaxError: expected err, got nil and %v max", max)
	}

	ch := sliceToChan(list)
	maxCh, err := MaxCh(ch)
	if err == nil {
		t.Errorf("TestMaxError MaxCh: expected err, got nil and %v max", maxCh)
	}
}
