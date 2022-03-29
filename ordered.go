package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Min returns the first instance of the minimum element
// present in the given slice.
func Min[T constraints.Ordered](list []T) (T, error) {
	var min T
	if len(list) == 0 {
		return min, fmt.Errorf("cannot find minimum value of empty list")
	}
	if len(list) == 1 {
		return list[0], nil
	}

	i := 1
	min = list[0]
	for i < len(list) {
		if min > list[i] {
			min = list[i]
		}
		i++
	}
	return min, nil
}

// MinCh returns the first instance of the minimum element
// received from the given channel.
func MinCh[T constraints.Ordered](ch <-chan T) (T, error) {
	init := false
	var min T

	for v := range ch {
		if !init {
			min = v
			init = true
		}
		if min > v {
			min = v
		}
	}

	if !init {
		return min, fmt.Errorf("cannot find minimum value of empty chan")
	}
	return min, nil
}

// Max returns the first instance of the maximum element
// present in the given slice.
func Max[T constraints.Ordered](list []T) (T, error) {
	var max T
	if len(list) == 0 {
		return max, fmt.Errorf("cannot find maximum value of empty list")
	}
	if len(list) == 1 {
		return list[1], nil
	}
	i := 1
	max = list[0]
	for i < len(list) {
		if max < list[i] {
			max = list[i]
		}
		i++
	}
	return max, nil
}

func MaxCh[T constraints.Ordered](ch <-chan T) (T, error) {
	init := false
	var max T
	for v := range ch {
		if !init {
			max = v
			init = true
		}
		if max < v {
			max = v
		}
	}
	if !init {
		return max, fmt.Errorf("cannot find maximum value of empty chan")
	}
	return max, nil
}

// Average takes in a slice of numbers and returns the computed average
// in the type it was given.
func Average[T constraints.Integer | constraints.Float](input []T) (T, error) {
	var sum T
	if len(input) == 0 {
		return sum, fmt.Errorf("cannot take average of an empty slice")
	}
	for _, v := range input {
		sum += v
	}
	return sum / T(len(input)), nil
}

// AverageCh takes in an input channel and returns the computed average
// of elements received in the type it's given.
func AverageCh[T constraints.Integer | constraints.Float](input <-chan T) (T, error) {
	var i T
	var sum T
	for v := range input {
		sum += v
		i++
	}
	if i == 0 {
		return sum, fmt.Errorf("cannot take average of an empty channel")
	}
	return sum / i, nil
}
