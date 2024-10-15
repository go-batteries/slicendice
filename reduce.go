package slicendice

// ReducerFunc is a generic function type that defines the operation to be
// applied to each element in a collection during reduction. It takes three
// arguments:
// - accumulator: the cumulative result (of type V) that is carried forward.
// - next: the next element (of type E) in the collection to be processed.
// - index: the index of the current element in the collection.
// It returns the updated accumulator (of type V).
type ReducerFunc[E, V any] func(accumulator V, next E, index int) V

// Reduce applies a reducer function to each element in a slice and accumulates
// the results. It takes three arguments:
//   - elements: a slice of elements of type E that you want to reduce.
//   - reducer: a function that defines how the reduction will proceed. It takes
//     the current accumulator value (V), the next element (E), and the index (int)
//     of the element being processed.
//   - accumulator: the initial value of the accumulator, of type V, which is
//     also the final result type after the reduction.
//
// The function returns the final accumulator value after processing all elements.
func Reduce[E, V any](elements []E, reducer ReducerFunc[E, V], accumulator V) V {
	if len(elements) == 0 {
		return accumulator
	}

	for i, el := range elements {
		accumulator = reducer(accumulator, el, i)
	}

	return accumulator
}
