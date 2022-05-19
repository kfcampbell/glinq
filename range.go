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
