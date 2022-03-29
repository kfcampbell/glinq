package main

import "fmt"

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
	for v := range ch {
		if v == elem {
			return i
		}
		i++
	}
	return -1
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

	for v := range ch {
		if v == elem {
			lastIndex = i
		}
		i++
	}
	return lastIndex
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
	for v := range ch {
		if is(v) {
			return v, nil
		}
	}
	var empty T
	return empty, fmt.Errorf("cannot find item in empty chan")
}
