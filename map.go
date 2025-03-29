package slicendice

import "reflect"

// MapperFunc is a generic function type that defines how each element in a
// collection is transformed. It takes two parameters:
// - element: the current element of type E being processed.
// - index: the index of the current element in the collection.
// It returns the transformed element of type V.
type MapperFunc[E, V any] func(element E, index int) V

// Map applies a mapping function to each element in a slice and returns a
// new slice of transformed elements. It takes three arguments:
// - elements: a slice of elements of type E to be transformed.
// - mapper: a function that defines how each element will be transformed.
// - It returns a new slice of type V containing the transformed elements.
func Map[E, V any](elements []E, mapper MapperFunc[E, V]) []V {
	if len(elements) == 0 {
		return []V{}
	}

	result := make([]V, len(elements))

	for i, el := range elements {
		result[i] = mapper(el, i)
	}

	return result
}

//  FlatMap

// Takes a flat input: []T
// Applies a callback on each element T, returning a nested structure
// Then flatten the result up to a given depth, producing a flat []V
//
// Since Go generics can't introspect arbitrary []...[]V, instead:
//
// The callback returns a custom iterator (like a generator)
// The iterator exposes a Next() (V, bool) or similar
// You walk the iterator and append results into the final slice
//
// The other way is to use reflection, but that destroys the type safety

func FlatMap[T, V any](input []T, callback func(T, int) FlatIterator[V]) []V {
	var result []V

	for i, item := range input {
		iter := callback(item, i)
		for {
			if v, ok := iter.Next(); ok {
				result = append(result, v)
			} else {
				break
			}
		}
	}

	return result
}

// FlatIterator is a channel based iterator
type FlatIterator[V any] interface {
	Next() (V, bool)
}

type ChanIterator[V any] struct {
	ch <-chan V
}

func (it *ChanIterator[V]) Next() (V, bool) {
	v, ok := <-it.ch
	return v, ok
}

func NewChanIterator[V any](input any, depth int) FlatIterator[V] {
	ch := make(chan V)

	go func() {
		defer close(ch)

		var walk func(any, int)
		walk = func(val any, remainingDepth int) {
			rv := reflect.ValueOf(val)
			if rv.Kind() == reflect.Slice && remainingDepth > 0 {
				for i := 0; i < rv.Len(); i++ {
					walk(rv.Index(i).Interface(), remainingDepth-1)
				}
			} else if v, ok := val.(V); ok {
				ch <- v
			}
		}

		walk(input, depth) // convention: bufferSize == depth
	}()

	return &ChanIterator[V]{ch}
}
