package main

// All takes in a slice and a predicate, and returns true
// if every element of the slice matches the predicate.
func All[T any](input []T, pred func(elem T) bool) bool {
	for _, v := range input {
		if !pred(v) {
			return false
		}
	}
	return true
}

// AllCh takes in a channel and a predicate, and returns true if every
// element received from the channel matches the predicate.
func AllCh[T any](input <-chan T, pred func(elem T) bool) bool {
	for v := range input {
		if !pred(v) {
			return false
		}
	}
	return true
}

// Any takes in a slice and a predicate, and returns true
// if any element of the slice matches the predicate.
func Any[T any](input []T, pred func(elem T) bool) bool {
	for _, v := range input {
		if pred(v) {
			return true
		}
	}
	return false
}

// AnyCh takes in a channel and a predicate, and returns true
// if any element received from the chan matches the predicate.
func AnyCh[T any](input <-chan T, pred func(elem T) bool) bool {
	for v := range input {
		if pred(v) {
			return true
		}
	}
	return false
}
