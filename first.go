package main

import (
	"fmt"
)

// IndexOf returns the index of the first instance of the given element
// in the given slice. If the given element is not present, -1 is returned.
func IndexOf[TSource comparable](source []TSource, value TSource) int {
	for i, v := range source {
		if v == value {
			return i
		}
	}

	return -1
}

// IndexOf returns the index of the first instance of the given element received
// from the given channel. If the given element is not present, -1 is returned.
func IndexOfCh[TSource comparable](source <-chan TSource, value TSource) int {
	i := 0

	for v := range source {
		if v == value {
			return i
		}
		i++
	}

	return -1
}

// LastIndexOf returns the index of the last instance of the given element in
// the given slice. If the given element is not present, -1 is returned.
func LastIndexOf[TSource comparable](source []TSource, value TSource) int {
	for i := len(source) - 1; i >= 0; i-- {
		if source[i] == value {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index of the last instance of the given element received
// from the given channel. If the given element is not present, -1 is returned.
func LastIndexOfCh[TSource comparable](source <-chan TSource, value TSource) int {
	i := 0
	lastIndex := -1

	for v := range source {
		if v == value {
			lastIndex = i
		}
		i++
	}

	return lastIndex
}

// First returns the first element in a slice that satisfies a specified condition,
// or an error if no element matches the specific condition.
func First[TSource any](source []TSource, predicate func(value TSource) bool) (TSource, error) {
	for _, v := range source {
		if predicate(v) {
			return v, nil
		}
	}

	var empty TSource
	return empty, fmt.Errorf("could not find item in input")
}

// FirstCh returns the first element received from a channel that satisfies a specified
// condition, or an error if no element matches the specific condition.
func FirstCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) (TSource, error) {
	for v := range source {
		if predicate(v) {
			return v, nil
		}
	}

	var empty TSource
	return empty, fmt.Errorf("cannot find item in chan")
}

// Last returns the last element in a slice that satisfies a specified condition,
// or an error if no element matches the specific condition.
func Last[TSource any](source []TSource, predicate func(value TSource) bool) (TSource, error) {
	for i := len(source) - 1; i >= 0; i-- {
		if predicate(source[i]) {
			return source[i], nil
		}
	}

	var empty TSource
	return empty, fmt.Errorf("could not find item in input")
}

// LastCh returns the last element received from a channel that satisfies a specified
// condition, or an error if no element matches the specific condition.
func LastCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) (TSource, error) {

	var empty TSource
	var doesExist = false
	for v := range source {
		if predicate(v) {
			empty = v
			doesExist = true
		}
	}
	if doesExist {
		return empty, nil
	}
	return empty, fmt.Errorf("cannot find item in chan")

}

// TODO: These do not exactly match Contains in LINQ, which instead would have this signature:
//   func Contains[TSource any](source []TSource, value TSource) bool

// Contains takes in a slice and a predicate, and returns true if the predicate
// matches any element in the slice.
func Contains[TSource any](source []TSource, predicate func(value TSource) bool) bool {
	_, err := First(source, predicate)
	return err == nil
}

// ContainsCh takes in a channel and a predicate, and returns true if the
// predicate matches any element sent down the channel.
func ContainsCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) bool {
	_, err := FirstCh(source, predicate)
	return err == nil
}

// Count returns the number of elements in the given slice.
func Count[TSource any](source []TSource) int {
	return len(source)
}

// CountCh takes in a channel and returns the number of elements
// passed down the channel.
func CountCh[TSource any](source <-chan TSource) int {
	i := 0
	for _ = range source {
		i++
	}
	return i
}
