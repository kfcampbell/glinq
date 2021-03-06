# glinq

glinq (prounced "glink") is an attempt to provide a LINQ API using Go generics. The name uncreatively comes from combining Go and LINQ. Inspired by github.com/samber/lo

Long-term TODOs:

- Implement parallel operations
- Use fuzzing in testing
- Full coverage of functions in LINQ
- Complete testing with code coverage
- Usage instructions and examples
- Cancellation (esp. in channel operations)
- Decide whether to implement container types (e.g. Set) like .NET has
- Use `t.Run(tc.name, ...)` as the testing pattern rather than the janky string constructions currently in place.
- Consider using literal instantiations instead of `make` everywhere
- Organize file structure better

[LINQ API scheme](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=net-6.0):

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
- [x] Count
- [ ] DefaultIfEmpty
  - not worth implementing?
- [x] Distinct
- [x] DistinctBy
- [ ] ElementAt
  - not worth implementing?
- [ ] ElementAtOrDefault
  - not worth implementing?
- [ ] Empty
  - not worth implementing?
- [x] Except
- [x] ExceptBy
- [x] First
- [ ] FirstOrDefault
  - not worth implementing?
- [ ] GroupBy
  - is this possible without anonymous types?
- [ ] GroupJoin
- [x] Intersect
- [x] IntersectBy
- [ ] Join
  - is this possible without anonymous types?
- [x] Last
- [ ] LastOrDefault
  - not worth implementing?
- [x] LongCount
- [x] Max
- [x] MaxBy
- [x] Min
- [x] MinBy
- [ ] OfType
  - is this possible with Go slices/channels?
- [x] OrderBy
- [x] OrderByDescending
- [x] Prepend
- [x] Range
- [x] Repeat
- [ ] Reverse
- [x] Select
- [ ] SelectMany
- [ ] SequenceEqual
- [ ] Single
- [ ] SingleOrDefault
- [x] Skip
- [x] SkipLast
- [x] SkipWhile
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
