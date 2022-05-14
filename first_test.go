package main

import "testing"

func TestIndexOfInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		value    int
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
		actIndexOf := IndexOf(tc.input, tc.value)
		if actIndexOf != tc.expected {
			t.Errorf("TestIndexOfInt %v: expected %v, got %v", tc.name, tc.value, actIndexOf)
		}

		ch := sliceToChan(tc.input)
		actIndexOfCh := IndexOfCh(ch, tc.value)
		if actIndexOfCh != tc.expected {
			t.Errorf("TestIndexOfInt IndexOfCh %v: expected %v, got %v", tc.name, tc.value, actIndexOfCh)
		}
	}
}

func TestIndexOfString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		value    string
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
		actIndexOf := IndexOf(tc.input, tc.value)
		if actIndexOf != tc.expected {
			t.Errorf("TestIndexOfString %v: expected %v, got %v", tc.name, tc.value, actIndexOf)
		}

		ch := sliceToChan(tc.input)
		actIndexOfCh := IndexOfCh(ch, tc.value)
		if actIndexOfCh != tc.expected {
			t.Errorf("TestIndexOfString IndexOfCh %v: expected %v, got %v", tc.name, tc.value, actIndexOfCh)
		}
	}
}

func TestLastIndexOfInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		value    int
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
		actLastIndexOf := LastIndexOf(tc.input, tc.value)
		if actLastIndexOf != tc.expected {
			t.Errorf("TestLastIndexOfInt %v: expected %v, got %v", tc.name, tc.value, actLastIndexOf)
		}

		ch := sliceToChan(tc.input)
		actLastIndexOfCh := LastIndexOfCh(ch, tc.value)
		if actLastIndexOfCh != tc.expected {
			t.Errorf("TestLastIndexOfInt LastIndexOfCh %v: expected %v, got %v", tc.name, tc.value, actLastIndexOfCh)
		}
	}
}

func TestLastIndexOfString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		value    string
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
		actLastIndexOf := LastIndexOf(tc.input, tc.value)
		if actLastIndexOf != tc.expected {
			t.Errorf("TestLastIndexOfString %v: expected %v, got %v", tc.name, tc.expected, actLastIndexOf)
		}

		ch := sliceToChan(tc.input)
		actLastIndexOfCh := LastIndexOfCh(ch, tc.value)
		if actLastIndexOfCh != tc.expected {
			t.Errorf("TestLastIndexOfString LastIndexOfCh %v: expected %v, got %v", tc.name, tc.expected, actLastIndexOfCh)
		}
	}
}

func TestFirstInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		pred     func(value int) bool
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			func(value int) bool {
				return value == 2
			},
			2,
		},
		{
			"expandedCase",
			[]int{1, 4, 3, 23, 412, 19, 49, 67, 76, 15, 23, 18},
			func(value int) bool {
				return value == 412
			},
			412,
		},
	}
	for _, tc := range cases {
		actFirst, err := First(tc.input, tc.pred)
		if err != nil || actFirst != tc.expected {
			t.Errorf("TestFirstInt %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirst, err)
		}

		ch := sliceToChan(tc.input)
		actFirstCh, err := FirstCh(ch, tc.pred)
		if err != nil || actFirstCh != tc.expected {
			t.Errorf("TestFirstInt FirstCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirstCh, err)
		}
	}
}

func TestFirstIntErr(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		pred  func(value int) bool
	}{
		{
			"doesNotExist",
			[]int{1, 4, 3},
			func(value int) bool {
				return value == 412
			},
		},
		{
			"emptySlice",
			[]int{},
			func(value int) bool {
				return value == 412
			},
		},
	}

	for _, tc := range cases {
		actFirst, err := First(tc.input, tc.pred)
		if err == nil {
			t.Errorf("TestFirstIntErr %v: wanted err but got nil, found :%v", tc.name, actFirst)
		}

		ch := sliceToChan(tc.input)
		actFirstCh, err := FirstCh(ch, tc.pred)
		if err == nil {
			t.Errorf("TestFirstIntErr FirstCh %v: wanted err but got nil, found :%v", tc.name, actFirstCh)
		}
	}
}

func TestFirstString(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		pred     func(value string) bool
		expected string
	}{
		{
			"happyCase",
			[]string{"abc", "def", "hijk"},
			func(value string) bool {
				return value == "def"
			},
			"def",
		},
		{
			"expandedCase",
			[]string{"abc", "def", "hijk", "lmnop", "qrs", "tuv", "wxy", "z"},
			func(value string) bool {
				return value == "hijk"
			},
			"hijk",
		},
	}
	for _, tc := range cases {
		actFirst, err := First(tc.input, tc.pred)
		if err != nil || actFirst != tc.expected {
			t.Errorf("TestFirstString %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirst, err)
		}

		ch := sliceToChan(tc.input)
		actFirstCh, err := FirstCh(ch, tc.pred)
		if err != nil || actFirstCh != tc.expected {
			t.Errorf("TestFirstString FirstCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirstCh, err)
		}
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		pred     func(value int) bool
		expected bool
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			func(value int) bool {
				return value == 2
			},
			true,
		},
		{
			"expandedCase",
			[]int{1, 4, 3, 23, 412, 19, 49, 67, 76, 15, 23, 18},
			func(value int) bool {
				return value == 412
			},
			true,
		},
		{
			"missingData",
			[]int{1, 2, 3},
			func(value int) bool {
				return value == 4
			},
			false,
		},
	}
	for _, tc := range cases {
		actContains := Contains(tc.input, tc.pred)
		if actContains != tc.expected {
			t.Errorf("TestContains %v: expected %v, got %v", tc.name, tc.expected, actContains)
		}

		ch := sliceToChan(tc.input)
		actContainsCh := ContainsCh(ch, tc.pred)
		if actContainsCh != tc.expected {
			t.Errorf("TestContains ContainsCh %v: expected %v, got %v", tc.name, tc.expected, actContainsCh)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			"emptyCase",
			[]int{},
			0,
		},
		{
			"expandedCase",
			[]int{1, 9, 2, 5, 45, 98, 26, 79, 100, 101, 432, 19, 8, 4, 16},
			15,
		},
	}

	for _, tc := range cases {
		actCount := Count(tc.input)
		if tc.expected != actCount {
			t.Errorf("TestCount %v: expected %v, got %v", tc.name, tc.expected, actCount)
		}

		ch := sliceToChan(tc.input)
		actCountCh := CountCh(ch)
		if tc.expected != actCountCh {
			t.Errorf("TestCount CountCh %v: expected %v, got %v", tc.name, tc.expected, actCountCh)
		}
	}
}

func TestLastInt(t *testing.T){
	cases := []struct {
		name     string
		input    []int
		pred     func(value int) bool
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3},
			func(value int) bool {
				return value == 2
			},
			2,
		},
		{
			"expandedCase",
			[]int{1, 4, 3, 23, 412, 19, 49, 67, 76, 15, 23, 18},
			func(value int) bool {
				return value == 412
			},
			412,
		},
	}
	for _, tc := range cases {
		actFirst, err := Last(tc.input, tc.pred)
		if err != nil || actFirst != tc.expected {
			t.Errorf("TestFirstInt %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirst, err)
		}

		ch := sliceToChan(tc.input)
		actFirstCh, err := LastCh(ch, tc.pred)
		if err != nil || actFirstCh != tc.expected {
			t.Errorf("TestFirstInt FirstCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actFirstCh, err)
		}
	}
}