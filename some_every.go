package slicendice

// PredicateFunc is a generic function type that defines the condition for
// determining if at least one or all elements in a collection meets the criteria.
// It takes two parameters:
// - element: the current element of type E being processed.
// - index: the index of the current element in the collection.
// It returns true if the element meets the criteria.
type PredicateFunc[E any] func(element E, index int) bool

// Some checks if at least one element in the slice meets the condition defined
// by the provided predicate function. It takes two arguments:
// - elements: a slice of elements of type E to be checked.
// - predicate: a function that defines the condition to check for each element.
// It returns true if at least one element satisfies the condition; otherwise, false.
func Some[E any](elements []E, predicate PredicateFunc[E]) bool {
	for i, el := range elements {
		if predicate(el, i) {
			return true
		}
	}
	return false
}

// Every checks if all elements in the slice meet the condition defined by the
// provided predicate function. It takes two arguments:
// - elements: a slice of elements of type E to be checked.
// - predicate: a function that defines the condition to check for each element.
// It returns true if all elements satisfy the condition; otherwise, false.
func Every[E any](elements []E, predicate PredicateFunc[E]) bool {
	for i, el := range elements {
		if !predicate(el, i) {
			return false
		}
	}
	return true
}
