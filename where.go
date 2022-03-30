package main

// Where filters a slice of values based on a predicate
func Where[TSource any](source []TSource, predicate func(value TSource) bool) []TSource {
	result := make([]TSource, 0)

	for _, v := range source {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// Where filters a channel of values based on a predicate
func WhereCh[TSource any](source <-chan TSource, predicate func(value TSource) bool) <-chan TSource {
	out := make(chan TSource)

	go func() {
		for v := range source {
			if predicate(v) {
				out <- v
			}
		}

		close(out)
	}()

	return out
}
