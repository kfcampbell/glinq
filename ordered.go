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
