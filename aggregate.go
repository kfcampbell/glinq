package main

func Aggregate[TSource, TResult any](input []TSource, seed TResult, agg func(seed TResult, elem TSource) TResult) TResult {
	for _, v := range input {
		seed = agg(seed, v)
	}
	return seed
}

func AggregateCh[TSource, TResult any](input <-chan TSource, seed TResult, agg func(seed TResult, elem TSource) TResult) TResult {
	for v := range input {
		seed = agg(seed, v)
	}
	return seed
}
