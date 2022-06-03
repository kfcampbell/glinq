package main

// Prepend adds a value to the beginning of the sequence
func Prepend[Tsource any](source []Tsource, value Tsource) []Tsource {
	return append([]Tsource{value}, source...)
}

// PrependCh adds a value to the beginning of the channel
func PrependCh[Tsource any](source <-chan Tsource, value Tsource) <-chan Tsource {
	result := make(chan Tsource)

	go func() {
		result <- value
		for v := range source {
			result <- v
		}

		close(result)
	}()

	return result
}
