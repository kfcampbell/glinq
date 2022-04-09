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
// If the source slice is empty, it returns the default value of TSource.
// TODO(kfcampbell): is this desirable behavior? should this return an error instead?
func MaxBy[TSource comparable, TKey constraints.Ordered](source []TSource, key func(elem TSource) TKey) TSource {
	var max TSource
	if len(source) == 0 {
		return max
	}
	if len(source) == 1 {
		return source[0]
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

	return max
}

// Returns the maximum value in a generic sequence according to a specified key selector function.
// If the source channel doesn't receive any values, it returns the default value of TSource.
// TODO(kfcampbell): is this desirable behavior? should this return an error instead?
func MaxByCh[TSource comparable, TKey constraints.Ordered](source <-chan TSource, key func(elem TSource) TKey) TSource {
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
	return max
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
