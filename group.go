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

func IntersectBy[TSource comparable, TKey comparable](first []TSource, second []TKey, key func(elem TSource) TKey) []TSource {
	result := make([]TSource, 0)
	firstSeen := make(map[TKey]TSource)
	secondSeen := make(map[TKey]struct{})
	for _, v := range first {
		elem := key(v)
		if _, ok := firstSeen[elem]; !ok {
			firstSeen[elem] = v
		}
	}

	for _, v := range second {
		_, existsInSecond := secondSeen[v]
		if !existsInSecond {
			secondSeen[v] = struct{}{}
		}
		firstElem, ok := firstSeen[v]
		if ok && !existsInSecond {
			result = append(result, firstElem)
		}
	}
	return result
}

func IntersectByCh[TSource comparable, TKey comparable](first <-chan TSource, second <-chan TKey, key func(elem TSource) TKey) <-chan TSource {
	results := make(chan TSource)
	firstSeen := make(map[TKey]TSource)
	secondSeen := make(map[TKey]struct{})
	go func() {
		for v := range first {
			firstKey := key(v)
			_, exists := firstSeen[firstKey]
			if !exists {
				firstSeen[firstKey] = v
			}
		}

		for v := range second {
			firstElem, existsInFirst := firstSeen[v]
			_, existsInSecond := secondSeen[v]
			if !existsInSecond {
				secondSeen[v] = struct{}{}
			}

			if existsInFirst && !existsInSecond {
				results <- firstElem
			}
		}
		close(results)
	}()
	return results
}
