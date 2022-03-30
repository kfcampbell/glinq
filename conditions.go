package main

// All determines whether all elements of a slice satisfy a condition.
func All[TSource any](source []TSource, predicate func(value TSource) bool) bool {
	for _, v := range source {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// AllCh determines whether all elements received from a channel satisfy a condition.
func AllCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) bool {
	for v := range source {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Any determines whether any element of a slice satisfies a condition.
func Any[TSource any](source []TSource, predicate func(value TSource) bool) bool {
	for _, v := range source {
		if predicate(v) {
			return true
		}
	}

	return false
}

// AnyCh determines whether any element received from a channel satisfies a condition.
func AnyCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) bool {
	for v := range source {
		if predicate(v) {
			return true
		}
	}

	return false
}
