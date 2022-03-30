package main

// Aggregate applies an accumulator function over a slice. The specified seed value is
// used as the initial accumulator value.
func Aggregate[TSource, TAccumulate any](source []TSource, seed TAccumulate, agg func(seed TAccumulate, value TSource) TAccumulate) TAccumulate {
	for _, v := range source {
		seed = agg(seed, v)
	}

	return seed
}

// AggregateCh applies an accumulator function over values received from a channel. The
// specified seed value is used as the initial accumulator value.
func AggregateCh[TSource, TAccumulate any](source <-chan TSource, seed TAccumulate, agg func(seed TAccumulate, value TSource) TAccumulate) TAccumulate {
	for v := range source {
		seed = agg(seed, v)
	}

	return seed
}
