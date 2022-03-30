package main

func chanToSlice[TSource any](source <-chan TSource) []TSource {
	result := make([]TSource, 0)

	for v := range source {
		result = append(result, v)
	}

	return result
}

func sliceToChan[TSource any](source []TSource) <-chan TSource {
	result := make(chan TSource)

	go func() {
		for _, v := range source {
			result <- v
		}

		close(result)
	}()

	return result
}

func sliceValueEquality[TSource comparable](left, right []TSource) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func sliceOfSliceValueEquality[TSource comparable](left, right [][]TSource) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if len(left[i]) != len(right[i]) {
			return false
		}

		for k := range left[i] {
			if left[i][k] != right[i][k] {
				return false
			}
		}
	}

	return true
}
