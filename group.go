package main

func Intersect[TSource comparable](first []TSource, second []TSource) []TSource {
	result := make([]TSource, 0)
	firstSeen := make(map[TSource]struct{})
	secondSeen := make(map[TSource]struct{})

	for _, v := range first {
		if _, ok := firstSeen[v]; !ok {
			firstSeen[v] = struct{}{}
		}
	}

	for _, v := range second {
		_, inSecond := secondSeen[v]
		if !inSecond {
			secondSeen[v] = struct{}{}
		}
		_, inFirst := firstSeen[v]
		if inFirst && !inSecond {
			result = append(result, v)
		}
	}
	return result
}

func IntersectCh[TSource comparable](first <-chan TSource, second <-chan TSource) <-chan TSource {
	result := make(chan TSource)
	firstSeen := make(map[TSource]struct{})
	secondSeen := make(map[TSource]struct{})

	go func() {

		for v := range first {
			if _, ok := firstSeen[v]; !ok {
				firstSeen[v] = struct{}{}
			}
		}

		for v := range second {
			_, inFirst := firstSeen[v]
			_, inSecond := secondSeen[v]
			if !inSecond {
				secondSeen[v] = struct{}{}
			}
			if inFirst && !inSecond {
				result <- v
			}
		}

		close(result)
	}()
	return result
}
