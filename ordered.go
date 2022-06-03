package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Min returns the minimum value in a slice of values
func Min[TSource constraints.Ordered](source []TSource) (TSource, error) {
	init := false
	var min TSource

	for _, v := range source {
		if !init {
			min = v
			init = true
		} else if min > v {
			min = v
		}
	}

	if !init {
		return min, fmt.Errorf("cannot find minimum value of empty slice")
	}

	return min, nil
}

// MinCh returns the minimum value in a channel of values
func MinCh[TSource constraints.Ordered](source <-chan TSource) (TSource, error) {
	init := false
	var min TSource

	for v := range source {
		if !init {
			min = v
			init = true
		} else if min > v {
			min = v
		}
	}

	if !init {
		return min, fmt.Errorf("cannot find minimum value of empty chan")
	}

	return min, nil
}

// MinBy returns the minimum value in a generic sequence according to a specified key selector function.
// If the given slice has no elements, an error is returned.
func MinBy[TSource comparable, TKey constraints.Ordered](source []TSource, key func(elem TSource) TKey) (TSource, error) {
	var min TSource
	if len(source) == 0 {
		return min, fmt.Errorf("cannot take minimum of empty slice")
	}
	if len(source) == 1 {
		return source[0], nil
	}

	minKey := key(source[0])
	min = source[0]
	for i := 1; i < len(source); i++ {
		key := key(source[i])
		if key < minKey {
			minKey = key
			min = source[i]
		}
	}
	return min, nil
}

// MinByCh returns the minimum value in a generic sequence according to a specified key selector function.
// If the given channel receives no elements, an error is returned.
func MinByCh[TSource comparable, TKey constraints.Ordered](source <-chan TSource, key func(elem TSource) TKey) (TSource, error) {
	var min TSource
	var minKey TKey
	first := false
	for v := range source {
		if !first {
			first = true
			min = v
			minKey = key(v)
		}
		key := key(v)
		if minKey > key {
			min = v
			minKey = key
		}
	}
	if !first {
		return min, fmt.Errorf("cannot take minimum of a chan with no values")
	}
	return min, nil
}

// Max returns the maximum value in a slice of values
func Max[TSource constraints.Ordered](source []TSource) (TSource, error) {
	init := false
	var max TSource

	for _, v := range source {
		if !init {
			max = v
			init = true
		} else if max < v {
			max = v
		}
	}

	if !init {
		return max, fmt.Errorf("cannot find maximum value of empty slice")
	}

	return max, nil
}

// MaxCh returns the maximum value in a channel of values
func MaxCh[TSource constraints.Ordered](source <-chan TSource) (TSource, error) {
	init := false
	var max TSource

	for v := range source {
		if !init {
			max = v
			init = true
		} else if max < v {
			max = v
		}
	}

	if !init {
		return max, fmt.Errorf("cannot find maximum value of empty chan")
	}

	return max, nil
}

// Returns the maximum value in a generic sequence according to a specified key selector function.
// If the source slice is empty, it returns an error
func MaxBy[TSource comparable, TKey constraints.Ordered](source []TSource, key func(elem TSource) TKey) (TSource, error) {
	var max TSource
	if len(source) == 0 {
		return max, fmt.Errorf("cannot take max of empty slice")
	}
	if len(source) == 1 {
		return source[0], nil
	}
	max = source[0]
	maxKey := key(source[0])
	for i := 1; i < len(source); i++ {
		key := key(source[i])
		if maxKey < key {
			maxKey = key
			max = source[i]
		}
	}

	return max, nil
}

// Returns the maximum value in a generic sequence according to a specified key selector function.
// If the source channel doesn't receive any values, it returns an error
func MaxByCh[TSource comparable, TKey constraints.Ordered](source <-chan TSource, key func(elem TSource) TKey) (TSource, error) {
	var max TSource
	var maxKey TKey
	first := false
	for v := range source {
		if !first {
			first = true
			max = v
			maxKey = key(v)
			continue
		}
		key := key(v)
		if maxKey < key {
			maxKey = key
			max = v
		}
	}
	if !first {
		return max, fmt.Errorf("cannot take max of chan with no values")
	}
	return max, nil
}

// Average computes the average of a slice of values
func Average[TSource constraints.Integer | constraints.Float](source []TSource) (TSource, error) {
	i := TSource(len(source))
	var sum TSource

	for _, v := range source {
		sum += v
	}

	if i == 0 {
		return sum, fmt.Errorf("cannot take average of an empty slice")
	}

	return sum / i, nil
}

// AverageCh computes the average of a channel of values
func AverageCh[TSource constraints.Integer | constraints.Float](source <-chan TSource) (TSource, error) {
	var i TSource
	var sum TSource

	for v := range source {
		sum += v
		i++
	}

	if i == 0 {
		return sum, fmt.Errorf("cannot take average of an empty channel")
	}

	return sum / i, nil
}

func OrderBy[TSource comparable, TKey constraints.Ordered](source []TSource, key func(elem TSource) TKey) []TSource {
	result := make([]TSource, 0)
	pairs := make(map[TKey]TSource)
	keys := make([]TKey, len(source))

	for i, v := range source {
		out := key(v)
		pairs[out] = v

		// TODO(kfcampbell): put key in sorted order here to prevent unnecessary
		// sorting below
		keys[i] = out
	}

	keys = quickSort(keys)

	// iterate through sorted keys and append to result slice in order
	for _, v := range keys {
		result = append(result, pairs[v])
	}
	return result
}

func OrderByCh[TSource comparable, TKey constraints.Ordered](source <-chan TSource, key func(elem TSource) TKey) <-chan TSource {
	result := make(chan TSource)

	go func() {
		pairs := make(map[TKey]TSource)
		keys := make([]TKey, 0)

		for v := range source {
			out := key(v)
			pairs[out] = v
			keys = append(keys, out)
		}

		keys = quickSort(keys)

		for _, v := range keys {
			result <- pairs[v]
		}
		close(result)
	}()

	return result
}

func quickSort[TSource constraints.Ordered](input []TSource) []TSource {
	for i := 1; i < len(input); i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j = j - 1
		}
	}
	return input
}
