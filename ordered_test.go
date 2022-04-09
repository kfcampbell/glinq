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
			name:   "happyCase",
			source: []int{6, 5, 4, 3, 2, 1},
			key: func(elem int) int {
				return elem * -1
			},
			expected: 1,
		},
		{
			name:   "emptySource",
			source: []int{},
			key: func(elem int) int {
				return elem * -1
			},
			expected: 0,
		},
	}

	for _, tc := range cases {
		actual := MaxBy(tc.source, tc.key)
		if actual != tc.expected {
			t.Errorf("TestMaxBy %v: expected %v, got %v", tc.name, tc.expected, actual)
		}

		ch := sliceToChan(tc.source)
		actualCh := MaxByCh(ch, tc.key)
		if actualCh != tc.expected {
			t.Errorf("TestMaxBy MaxByCh %v: expected %v, got %v", tc.name, tc.expected, actualCh)
		}
	}
}
