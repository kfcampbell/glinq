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
