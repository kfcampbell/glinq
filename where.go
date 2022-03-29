package main

// Where takes in a slice and a predicate, and returns all elements
// of the given slice for which the predicate is true.
func Where[T any](input []T, pred func(elem T) bool) []T {
	res := make([]T, 0)
	for _, v := range input {
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}

// WhereCh takes in a channel and a predicate, and returns a channel which will
// receive all elements for the given channel for which the predicate is true.
func WhereCh[T any](in <-chan T, pred func(elem T) bool) <-chan T {
	out := make(chan T)

	go func() {
		for v := range in {
			if pred(v) {
				out <- v
			}
		}

		close(out)
	}()

	return out
}
