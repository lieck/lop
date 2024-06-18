package loplus

import (
	"github.com/samber/lo"
)

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V) bool) []V {
	return lo.Filter(collection, func(item V, _ int) bool {
		return predicate(item)
	})
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	return lo.Map(collection, func(item T, _ int) R {
		return iteratee(item)
	})
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(item T) (R, bool)) []R {
	return lo.FilterMap(collection, func(item T, _ int) (R, bool) {
		return callback(item)
	})
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
func FlatMap[T any, R any](collection []T, iteratee func(item T) []R) []R {
	return lo.FlatMap(collection, func(item T, _ int) []R {
		return iteratee(item)
	})
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	return lo.Reduce(collection, func(agg R, item T, _ int) R {
		return accumulator(agg, item)
	}, initial)
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	return lo.ReduceRight(collection, func(agg R, item T, _ int) R {
		return accumulator(agg, item)
	}, initial)
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(item T)) {
	lo.ForEach(collection, func(item T, _ int) {
		iteratee(item)
	})
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// Play: https://go.dev/play/p/vgQj3Glr6lT
func Times[T any](count int, iteratee func(index int) T) []T {
	return lo.Times(count, iteratee)
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
// Play: https://go.dev/play/p/DTzbeXZ6iEN
func Uniq[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
// Play: https://go.dev/play/p/g42Z3QSb53u
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return lo.UniqBy(collection, iteratee)
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// Play: https://go.dev/play/p/XnQBd_v6brd
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(collection, iteratee)
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// Play: https://go.dev/play/p/NfQ_nGjkgXW
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	return lo.PartitionBy(collection, iteratee)
}

// Flatten returns an array a single level deep.
// Play: https://go.dev/play/p/rbp9ORaMpjw
func Flatten[T any](collection [][]T) []T {
	return lo.Flatten(collection)
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/Qp73bnTDnc7
func Shuffle[T any](collection []T) []T {
	return lo.Shuffle(collection)
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/fhUMLvZ7vS6
func Reverse[T any](collection []T) []T {
	return lo.Reverse(collection)
}

// TODO Repeat

// RepeatBy builds a slice with values returned by N calls of callback.
// Play: https://go.dev/play/p/ozZLCtX_hNU
func RepeatBy[T any](count int, predicate func(index int) T) []T {
	return lo.RepeatBy(count, predicate)
}

// Chunk returns an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
// Play: https://go.dev/play/p/EeKl0AuTehH
func Chunk[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

// KeyBy transforms a slice or an array of structs to a map based on a pivot callback.
// Play: https://go.dev/play/p/mdaClUAT-zZ
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	return lo.KeyBy(collection, iteratee)
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return lo.Associate(collection, transform)
}

// SliceToMap returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of Associate().
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return lo.SliceToMap(collection, transform)
}

// Drop drops n elements from the beginning of a slice or array.
// Play: https://go.dev/play/p/JswS7vXRJP2
func Drop[T any](collection []T, n int) []T {
	return lo.Drop(collection, n)
}

// DropRight drops n elements from the end of a slice or array.
// Play: https://go.dev/play/p/GG0nXkSJJa3
func DropRight[T any](collection []T, n int) []T {
	return lo.DropRight(collection, n)
}

// DropWhile drops elements from the beginning of a slice or array while the predicate returns true.
// Play: https://go.dev/play/p/7gBPYw2IK16
func DropWhile[T any](collection []T, predicate func(item T) bool) []T {
	return lo.DropWhile(collection, predicate)
}

// DropRightWhile drops elements from the end of a slice or array while the predicate returns true.
// Play: https://go.dev/play/p/3-n71oEC0Hz
func DropRightWhile[T any](collection []T, predicate func(item T) bool) []T {
	return lo.DropRightWhile(collection, predicate)
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
func Reject[V any](collection []V, predicate func(item V) bool) []V {
	return lo.Reject(collection, func(item V, _ int) bool {
		return predicate(item)
	})
}

// Count counts the number of elements in the collection that compare equal to value.
// Play: https://go.dev/play/p/Y3FlK54yveC
func Count[T comparable](collection []T, value T) (count int) {
	return lo.Count(collection, value)
}

// CountBy counts the number of elements in the collection for which predicate is true.
// Play: https://go.dev/play/p/ByQbNYQQi4X
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	return lo.CountBy(collection, predicate)
}

// CountValues counts the number of each element in the collection.
// Play: https://go.dev/play/p/-p-PyLT4dfy
func CountValues[T comparable](collection []T) map[T]int {
	return lo.CountValues(collection)
}

// CountValuesBy counts the number of each element return from mapper function.
// Is equivalent to chaining lo.Map and lo.CountValues.
// Play: https://go.dev/play/p/2U0dG1SnOmS
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) U) map[U]int {
	return lo.CountValuesBy(collection, mapper)
}

// Subset returns a copy of a slice from `offset` up to `length` elements. Like `slice[start:start+length]`, but does not panic on overflow.
// Play: https://go.dev/play/p/tOQu1GhFcog
func Subset[T any](collection []T, offset int, length uint) []T {
	return lo.Subset(collection, offset, length)
}

// Slice returns a copy of a slice from `start` up to, but not including `end`. Like `slice[start:end]`, but does not panic on overflow.
// Play: https://go.dev/play/p/8XWYhfMMA1h
func Slice[T any](collection []T, start int, end int) []T {
	return lo.Slice(collection, start, end)
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/XfPzmf9gql6
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	return lo.Replace(collection, old, new, n)
}
