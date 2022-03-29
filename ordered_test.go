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
