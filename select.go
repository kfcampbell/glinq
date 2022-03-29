package main

// Select takes in an input slice and a selector function, and returns
// a slice of the selected TResult. This operation is similar to Map in
// other ecosystems.
func Select[TSource, TResult any](input []TSource, selector func(in TSource) TResult) []TResult {
	res := make([]TResult, 0)
	for _, elem := range input {
		sel := selector(elem)
		res = append(res, sel)
	}
	return res
}

// SelectCh takes in an input channel and a selector function, and returns
// a channel of the selected result.
func SelectCh[TSource, TResult any](in <-chan TSource, selector func(in TSource) TResult) <-chan TResult {
	out := make(chan TResult)

	go func() {
		for v := range in {
			elem := selector(v)
			out <- elem
		}
		close(out)
	}()

	return out
}
