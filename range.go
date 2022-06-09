package main

func Range(start, end int) []int {
	result := make([]int, end-start)
	for i := start; i < end; i++ {
		result[i-start] = i
	}
	return result
}

func Repeat[Tsource any](elem Tsource, repeat int) []Tsource {
	result := make([]Tsource, repeat)
	for i := 0; i < repeat; i++ {
		result[i] = elem
	}
	return result
}

func Skip[Tsource any](source []Tsource, skip int) []Tsource {
	result := make([]Tsource, len(source)-skip)
	for i := 0; i < len(source)-skip; i++ {
		result[i] = source[i+skip]
	}
	return result
}

func SkipCh[Tsource any](sourceCh <-chan Tsource, skip int) <-chan Tsource {
	result := make(chan Tsource)
	go func() {
		for i := 0; i < skip; i++ {
			<-sourceCh
		}
		for v := range sourceCh {
			result <- v
		}
		close(result)
	}()
	return result
}

func SkipLast[Tsource any](source []Tsource, count int) []Tsource {
	result := make([]Tsource, len(source)-count)
	for i := 0; i < len(source)-count; i++ {
		result[i] = source[i]
	}
	return result
}

func SkipWhile[Tsource any](source []Tsource, predicate func(value Tsource) bool) []Tsource {
	result := make([]Tsource, 0)
	for _, v := range source {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func SkipWhileCh[Tsource any](source <-chan Tsource, predicate func(value Tsource) bool) <-chan Tsource {
	result := make(chan Tsource)
	go func() {
		for v := range source {
			if !predicate(v) {
				result <- v
			}
		}
		close(result)
	}()
	return result
}
