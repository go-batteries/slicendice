package slicendice

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
