package main

import (
	"testing"
)

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
	input := make([]int, 0)

	min, err := Min(input)
	if err == nil {
		t.Errorf("TestMinError: expected err, got nil and %v min", min)
	}

	ch := sliceToChan(input)
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
	input := make([]int, 0)

	max, err := Max(input)
	if err == nil {
		t.Errorf("TestMaxError: expected err, got nil and %v max", max)
	}

	ch := sliceToChan(input)
	maxCh, err := MaxCh(ch)
	if err == nil {
		t.Errorf("TestMaxError MaxCh: expected err, got nil and %v max", maxCh)
	}
}

func TestAverageInt(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"happyCase",
			[]int{3, 4, 5},
			4,
		},
		{
			"truncate",
			[]int{3, 4},
			3,
		},
	}

	for _, tc := range cases {
		actualAvg, err := Average(tc.input)
		if actualAvg != tc.expected || err != nil {
			t.Errorf("TestAverageInt %v: expected %v, got %v, err %v", tc.name, tc.expected, actualAvg, err)
		}

		ch := sliceToChan(tc.input)
		actualAvgCh, err := AverageCh(ch)
		if actualAvgCh != tc.expected || err != nil {
			t.Errorf("TestAverageInt AverageCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actualAvgCh, err)
		}
	}
}

func TestAverageIntErr(t *testing.T) {
	input := make([]int, 0)

	actualAvg, err := Average(input)
	if err == nil {
		t.Errorf("TestAverageIntErr wanted err, got nil, return %v", actualAvg)
	}

	ch := sliceToChan(input)
	actualAvgCh, err := AverageCh(ch)
	if err == nil {
		t.Errorf("TestAverageIntErr AverageCh wanted err, got nil, return %v", actualAvgCh)
	}
}

func TestAverageFloat(t *testing.T) {
	cases := []struct {
		name     string
		input    []float32
		expected float32
	}{
		{
			"happyCase",
			[]float32{3.5, 4.0, 6.0},
			4.5,
		},
		{
			"repeating",
			[]float32{3.5, 4.0, 5.5},
			4.3333335,
		},
	}

	for _, tc := range cases {
		actualAvg, err := Average(tc.input)
		if actualAvg != tc.expected || err != nil {
			t.Errorf("TestAverageInt %v: expected %v, got %v, err %v", tc.name, tc.expected, actualAvg, err)
		}

		ch := sliceToChan(tc.input)
		actualAvgCh, err := AverageCh(ch)
		if actualAvgCh != tc.expected || err != nil {
			t.Errorf("TestAverageInt AverageCh %v: expected %v, got %v, err %v", tc.name, tc.expected, actualAvgCh, err)
		}
	}
}

func TestMaxByInt(t *testing.T) {
	cases := []struct {
		name     string
		source   []int
		key      func(elem int) int
		expected int
	}{
		{
			"happyCase",
			[]int{6, 5, 4, 3, 2, 1},
			func(elem int) int {
				return elem * -1
			},
			1,
		},
	}

	for _, tc := range cases {
		actual, err := MaxBy(tc.source, tc.key)
		if err != nil {
			t.Errorf("TestMaxBy %v: expected %v, got %v", tc.name, tc.expected, err)
		}
		if actual != tc.expected {
			t.Errorf("TestMaxBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.source)
		actualCh, err := MaxByCh(ch, tc.key)
		if err != nil {
			t.Errorf("TestMaxBy MaxByCh %v: expected %v, got %v", tc.name, tc.expected, err)
		}
		if actualCh != tc.expected {
			t.Errorf("TestMaxBy MaxByCh %v: expected %v, got %v", tc.name, tc.expected, actualCh)
		}
	}
}

func TestMaxByError(t *testing.T) {
	input := make([]int, 0)
	keyFunc := func(elem int) int {
		return elem * -1
	}
	max, err := MaxBy(input, keyFunc)
	if err == nil {
		t.Errorf("TestMaxByError: expected err, got %v", max)
	}

	inputCh := sliceToChan(input)
	maxCh, err := MaxByCh(inputCh, keyFunc)
	if err == nil {
		t.Errorf("TestMaxByError TestMaxCh: expected err, got %v", maxCh)
	}
}

func TestMinByInt(t *testing.T) {
	cases := []struct {
		name     string
		source   []int
		key      func(elem int) int
		expected int
	}{
		{
			"happyCase",
			[]int{1, 2, 3, 4, 5, 6},
			func(elem int) int {
				return elem * -1
			},
			6,
		},
	}

	for _, tc := range cases {
		actual, err := MinBy(tc.source, tc.key)
		if err != nil {
			t.Errorf("TestMinBy %v: expected %v, got %v", tc.name, tc.expected, err)
		}
		if actual != tc.expected {
			t.Errorf("TestMinBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.source)
		actualCh, err := MinByCh(ch, tc.key)
		if err != nil {
			t.Errorf("TestMinBy MinByCh %v: expected %v, got %v", tc.name, tc.expected, err)
		}
		if actualCh != tc.expected {
			t.Errorf("TestMinBy MinByCh %v: expected %v, got %v", tc.name, tc.expected, actualCh)
		}
	}
}

func TestMinByError(t *testing.T) {
	input := make([]int, 0)
	keyFunc := func(elem int) int {
		return elem * -1
	}
	min, err := MinBy(input, keyFunc)
	if err == nil {
		t.Errorf("TestMinByError: expected err, got %v", min)
	}

	inputCh := sliceToChan(input)
	minCh, err := MaxByCh(inputCh, keyFunc)
	if err == nil {
		t.Errorf("TestMinByError TestMinCh: expected err, got %v", minCh)
	}
}

func TestOrderBy(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		key      func(elem int) int
		expected []int
	}{
		{
			"identityCase",
			[]int{1, 2, 3},
			func(elem int) int {
				return elem * 1
			},
			[]int{1, 2, 3},
		},
		{
			"otherCase",
			[]int{1, 2, 3},
			func(elem int) int {
				return elem*2 - elem*elem
			},
			[]int{3, 2, 1},
		},
	}

	for _, tc := range cases {
		actual := OrderBy(tc.input, tc.key)
		if !sliceValueEquality(tc.expected, actual) {
			t.Errorf("TestOrderBy: %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.expected)
		actualCh := OrderByCh(ch, tc.key)
		actualSl := chanToSlice(actualCh)
		if !sliceValueEquality(tc.expected, actualSl) {
			t.Errorf("TestOrderBy OrderByCh: %v: expected %v, got %v", tc.name, tc.expected, actualSl)
		}
	}
}

func TestOrderByDescending(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		key      func(elem int) int
		expected []int
	}{
		{
			"identityCase",
			[]int{1, 2, 3},
			func(elem int) int {
				return elem * 1
			},
			[]int{3, 2, 1},
		},
		{
			"otherCase",
			[]int{1, 2, 3},
			func(elem int) int {
				return elem*2 - elem*elem
			},
			[]int{1, 2, 3},
		},
	}

	for _, tc := range cases {
		actual := OrderByDescending(tc.input, tc.key)
		if !sliceValueEquality(tc.expected, actual) {
			t.Errorf("TestOrderByDescending: %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.expected)
		actualCh := OrderByDescendingCh(ch, tc.key)
		actualSl := chanToSlice(actualCh)
		if !sliceValueEquality(tc.expected, actualSl) {
			t.Errorf("TestOrderByDescending OrderByDescendingCh: %v: expected %v, got %v", tc.name, tc.expected, actualSl)
		}
	}
}

func TestQuicksort(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"happyCase",
			[]int{5, 3, 2, 4, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"longerCase",
			[]int{23, 32, 19, 18, 13, 29, 25, -1, 4, 0, -13},
			[]int{-13, -1, 0, 4, 13, 18, 19, 23, 25, 29, 32},
		},
	}

	for _, tc := range cases {
		actual := quickSort(tc.input)
		if !sliceValueEquality(tc.expected, actual) {
			t.Errorf("TestQuicksort %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestQuicksortDescending(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"happyCase",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"longerCase",
			[]int{23, 32, 19, 18, 13, 29, 25, -1, 4, 0, -13},
			[]int{32, 29, 25, 23, 19, 18, 13, 4, 0, -1, -13},
		},
	}

	for _, tc := range cases {
		actual := quickSortDescending(tc.input)
		if !sliceValueEquality(tc.expected, actual) {
			t.Errorf("TestQuicksort %v: expected %v, got %v", tc.name, tc.expected, actual)
		}
	}
}
