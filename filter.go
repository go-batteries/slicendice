package slicendice

// FilterFunc is a generic function type that defines the condition for
// including an element in the filtered result. It takes two parameters:
// - element: the current element of type E being processed.
// - index: the index of the current element in the collection.
// It returns true if the element should be included in the filtered result.
type FilterFunc[E any] func(element E, index int) bool

// Filter applies a filter function to each element in a slice and returns a
// new slice of elements that meet the condition. It takes two arguments:
// - elements: a slice of elements of type E to be filtered.
// - filter: a function that defines the condition for including elements.
// - It returns a new slice containing the filtered elements.
func Filter[E any](elements []E, filter FilterFunc[E]) []E {
	if len(elements) == 0 {
		return []E{}
	}

	var res []E

	for i, el := range elements {
		if filter(el, i) {
			res = append(res, el)
		}
	}
	return res
}
