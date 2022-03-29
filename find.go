package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// IndexOf returns the index of the first instance of the given element
// in the given slice. If the given element is not present, -1 is returned.
func IndexOf[T comparable](list []T, elem T) int {
	for i, v := range list {
		if v == elem {
			return i
		}
	}
	return -1
}

func IndexOfCh[T comparable](ch <-chan T, elem T) int {
	i := 0
	for {
		v, ok := <-ch
		if !ok {
			return -1
		}
		if v == elem {
			return i
		}
		i++
	}
}

// LastIndexOf returns the index of the last instance of the given element in
// the given slice. If the given element is not present, -1 is returned.
func LastIndexOf[T comparable](list []T, elem T) int {
	i := len(list) - 1
	for i >= 0 {
		if list[i] == elem {
			return i
		}
		i--
	}
	return -1
}

func LastIndexOfCh[T comparable](ch <-chan T, elem T) int {
	i := 0
	lastIndex := -1
	for {
		v, ok := <-ch
		if !ok {
			return lastIndex
		}
		if v == elem {
			lastIndex = i
		}
		i++
	}
}

// Find returns the element present and an error if the item is not present
func Find[T comparable](list []T, is func(a T) bool) (T, error) {
	for _, elem := range list {
		if is(elem) {
			return elem, nil
		}
	}
	var empty T
	return empty, fmt.Errorf("could not find item in list")
}

func FindCh[T comparable](ch <-chan T, is func(a T) bool) (T, error) {
	for {
		elem, ok := <-ch
		if !ok {
			return elem, fmt.Errorf("cannot find item in empty chan")
		}
		if is(elem) {
			return elem, nil
		}
	}
}

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
	for {
		v, ok := <-ch
		if !ok {
			break
		}
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
	for {
		v, ok := <-ch
		if !ok {
			break
		}
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
