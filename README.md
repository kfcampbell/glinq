# glinq

glinq (prounced "glink") is an attempt to provide a LINQ API using Go generics. The name uncreatively comes from combining Go and LINQ. Inspired by github.com/samber/lo

Long-term TODOs:
- Implement parallel operations
- Use fuzzing in testing
- Full coverage of functions in samber/lo
- Complete testing with code coverage
- Cancellation (esp. in channel operations)
- Decide whether to implement container type (e.g. Set) like .NET has
- Use `t.Run(tc.name, ...)` as the testing pattern rather than the janky string constructions currently in place.
- Consider using literal instantiations instead of `make` everywhere

LINQ [API scheme](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=net-6.0)?
- [x] Aggregate
- [x] All
- [x] Any
- [x] Append
    - won't implement
- [x] AsEnumerable
    - won't implement
- [x] Average
- [x] Cast
    - won't implement
- [x] Chunk
- [x] Concat
    - won't implement
- [x] Contains
- [ ] Count
- [ ] DefaultIfEmpty
- [ ] Distinct
- [ ] DistinctBy
- [ ] ElementAt
- [ ] ElementAtOrDefault
- [ ] Empty
- [ ] Except
- [ ] ExceptBy
- [ ] First
- [ ] FirstOrDefault
- [ ] GroupBy
- [ ] GroupJoin
- [ ] Intersect
- [ ] IntersectBy
- [ ] Join
- [ ] Last
- [ ] LastOrDefault
- [ ] LongCount
- [x] Max
- [ ] MaxBy
- [x] Min
- [ ] MinBy
- [ ] OfType
- [ ] OrderBy
- [ ] OrderByDescending
- [ ] Prepend
- [ ] Range
- [ ] Repeat
- [ ] Reverse
- [x] Select
- [ ] SelectMany
- [ ] SequenceEqual
- [ ] Single
- [ ] SingleOrDefault
- [ ] Skip
- [ ] SkipLast
- [ ] SkipWhile
- [ ] Sum
- [ ] Take
- [ ] TakeLast
- [ ] TakeWhile
- [ ] ThenBy
- [ ] ThenByDescending
- [ ] ToArray
- [ ] ToDictionary
- [ ] ToHashSet
- [ ] ToList
- [ ] ToLookup
- [ ] TryGetNonEnumeratedCount
- [ ] Union
- [ ] UnionBy
- [x] Where
- [ ] Zip

Is there a way to add these methods onto generic slices/chans? Probably not, but if possible, should we?