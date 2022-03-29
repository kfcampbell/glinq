package main

import (
	"testing"
)

func TestAggregate(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		seed     int
		agg      func(seed int, elem string) int
		expected int
	}{
		{
			"accumulatingLength",
			[]string{"one", "two", "three"},
			0,
			func(seed int, elem string) int {
				return seed + len(elem)
			},
			11,
		},
		{
			"emptyCase",
			[]string{},
			0,
			func(seed int, elem string) int {
				return seed + len(elem)
			},
			0,
		},
	}

	for _, tc := range cases {
		actualAgg := Aggregate(tc.input, tc.seed, tc.agg)
		if actualAgg != tc.expected {
			t.Errorf("TestAggregate %v: expected %v, got %v", tc.name, tc.expected, actualAgg)
		}

		ch := sliceToChan(tc.input)
		actualAggCh := AggregateCh(ch, tc.seed, tc.agg)
		if actualAggCh != tc.expected {
			t.Errorf("TestAggregate AggregateCh %v: expected %v, got %v", tc.name, tc.expected, actualAggCh)
		}
	}
}
