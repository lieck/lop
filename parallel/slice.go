package parallel

import lop "github.com/samber/lo/parallel"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	return lop.Map(collection, func(item T, _ int) R {
		return iteratee(item)
	})
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(item T)) {
	lop.ForEach(collection, func(item T, _ int) {
		iteratee(item)
	})
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is call in parallel.
func Times[T any](count int, iteratee func(index int) T) []T {
	return lop.Times(count, iteratee)
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lop.GroupBy(collection, iteratee)
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	return lop.PartitionBy(collection, iteratee)
}
