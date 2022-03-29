package main

func Where[T any](input []T, pred func(elem T) bool) []T {
	res := make([]T, 0)
	for _, v := range input {
		if pred(v) {
			res = append(res, v)
		}
	}
	return res
}

func WhereCh[T any](in <-chan T, pred func(elem T) bool) <-chan T {
	out := make(chan T)

	go func() {
		for {
			v, ok := <-in
			if !ok {
				break
			}
			if pred(v) {
				out <- v
			}
		}
		close(out)
	}()

	return out
}
