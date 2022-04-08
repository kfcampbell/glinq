package main

// Select projects each element of a slice into a new form
func Select[TSource, TResult any](source []TSource, selector func(value TSource) TResult) []TResult {
	result := make([]TResult, 0)

	for _, v := range source {
		sel := selector(v)
		result = append(result, sel)
	}

	return result
}

// SelectCh projects each element of a channel into a new form
func SelectCh[TSource, TResult any](source <-chan TSource, selector func(value TSource) TResult) <-chan TResult {
	result := make(chan TResult)

	go func() {
		for v := range source {
			value := selector(v)
			result <- value
		}

		close(result)
	}()

	return result
}

// Chunk splits the elements of a slice into chunks of size at most "size"
func Chunk[TSource any](source []TSource, size int) [][]TSource {
	result := make([][]TSource, 0, (len(source)/size)+1)
	chunk := make([]TSource, 0, size)

	for _, v := range source {
		if len(chunk) == size {
			result = append(result, chunk)
			chunk = make([]TSource, 0, size)
		}

		chunk = append(chunk, v)
	}

	if len(chunk) != 0 {
		result = append(result, chunk)
	}

	return result
}

// ChunkCh splits the elements of a channel into chunks of size at most "size"
func ChunkCh[TSource any](source <-chan TSource, size int) <-chan []TSource {
	result := make(chan []TSource)

	go func() {
		chunk := make([]TSource, 0, size)

		for v := range source {
			if len(chunk) == size {
				result <- chunk
				chunk = make([]TSource, 0, size)
			}
			chunk = append(chunk, v)
		}

		if len(chunk) != 0 {
			result <- chunk
		}

		close(result)
	}()

	return result
}

// Distinct returns a slice with all duplicate elements of the given slice removed.
func Distinct[TSource comparable](source []TSource) []TSource {
	seen := make(map[TSource]struct{})
	output := make([]TSource, 0)
	for _, v := range source {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}
	return output
}

// DistinctCh returns a channel that will receive all the distinct values
// from the given channel.
func DistinctCh[TSource comparable](source <-chan TSource) <-chan TSource {
	output := make(chan TSource)
	seen := make(map[TSource]struct{})
	go func() {
		for v := range source {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				output <- v
			}
		}
		close(output)
	}()

	return output
}

// DistinctBy applies the given key function to extract distinct elements from the given slice.
func DistinctBy[TSource, TResult comparable](source []TSource, key func(elem TSource) TResult) []TSource {
	seen := make(map[TResult]struct{})
	output := make([]TSource, 0)
	for _, v := range source {
		elem := key(v)
		if _, ok := seen[elem]; !ok {
			seen[elem] = struct{}{}
			output = append(output, v)
		}
	}
	return output
}

// DistinctByCh applies the given key function to extract distinct elements received from
// the given channel and passes them down the returned channel.
func DistinctByCh[TSource, TResult comparable](source <-chan TSource, key func(elem TSource) TResult) <-chan TSource {
	output := make(chan TSource)
	seen := make(map[TResult]struct{})

	go func() {
		for v := range source {
			elem := key(v)
			if _, ok := seen[elem]; !ok {
				seen[elem] = struct{}{}
				output <- v
			}
		}
		close(output)
	}()

	return output
}

// Except returns only the _unique_ elements in first that are not
// present in second
// TODO(kfcampbell): improve this heinous time complexity
func Except[TSource comparable](first []TSource, second []TSource) []TSource {
	result := make([]TSource, 0)
	distinct := Distinct(first)
	for _, v := range distinct {
		notInSecond := !Contains(second, func(value TSource) bool {
			return value == v
		})
		if notInSecond {
			result = append(result, v)
		}
	}
	return result
}

// ExceptCh returns only the unique elements in first that are not present in second
func ExceptCh[TSource comparable](first <-chan TSource, second <-chan TSource) <-chan TSource {
	result := make(chan TSource)
	seen := make(map[TSource]struct{})

	go func() {
		for v := range second {
			seen[v] = struct{}{}
		}

		distinct := DistinctCh(first)
		for v := range distinct {
			if _, ok := seen[v]; !ok {
				result <- v
			}
		}
		close(result)
	}()
	return result
}
