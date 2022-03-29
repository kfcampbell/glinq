# glinq

glinq (prounced "glink") is an attempt to provide a LINQ API using Go generics. The name uncreatively comes from combining Go and LINQ. Inspired by github.com/samber/lo

Long-term TODOs:
- Implement parallel operations
- Use fuzzing in testing
- Full coverage of functions in samber/lo
- Complete testing with code coverage
- Cancellation (esp. in channel operations)


LINQ vs. Lodash API naming/scheme?
	- try to match [API scheme here](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=net-6.0)
	- select
		- selectMany
	- where
	- aggregate
