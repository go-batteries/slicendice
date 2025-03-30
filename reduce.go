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

// Find: the first element matching the criteria
func Find[E any](elements []E, finder func(e E, i int) bool) (E, int) {
	var v E

	for i, el := range elements {
		if ok := finder(el, i); ok {
			return el, i
		}
	}

	return v, -1
}

func FindAll[E any](elements []E, finder func(e E, i int) bool) ([]E, bool) {
	v := []E{}

	for i, el := range elements {
		if ok := finder(el, i); ok {
			v = append(v, el)
		}
	}

	return v, len(v) > 0
}

// MapFilterFunc: receives each item from an collection of elements
// and returns a transformed value and wether to include the result
// It takes two parameters:
// - element: the current element of type E being processed.
// - index: the index of the current element in the collection.
// It returns the transformed element of type V and if transformation was skipped
type MapFilterFunc[E, V any] func(element E, index int) (V, bool)

// MapFilter: combines a map and a filter
func MapFilter[E, V any](elements []E, mapfilter MapFilterFunc[E, V]) []V {
	if len(elements) == 0 {
		return []V{}
	}

	results := []V{}

	for i, el := range elements {
		if res, ok := mapfilter(el, i); ok {
			results = append(results, res)
		}
	}

	return results
}
