package lazy

import "github.com/go-batteries/slicendice"

func Map[E, V any](elements []E, mapper slicendice.MapperFunc[E, V]) Iter[V] {
	iter := ToSliceIter(elements)

	return NewFuncIterator(iter, func(e E, _ int) (bool, V) {
		return true, mapper(e, iter.index)
	})

}

func MapFromIter[E, V any](iter Iter[E], mapper slicendice.MapperFunc[E, V]) Iter[V] {
	return NewFuncIterator(iter, func(e E, i int) (bool, V) {
		return true, mapper(e, i)
	})

}

func Filter[E any](elements []E, mapper slicendice.FilterFunc[E]) Iter[E] {
	iter := ToSliceIter(elements)

	return NewFuncIterator(iter, func(e E, _ int) (bool, E) {
		return mapper(e, iter.index), e
	})

}

func FilterFromIter[E any](iter Iter[E], mapper slicendice.FilterFunc[E]) Iter[E] {
	return NewFuncIterator(iter, func(e E, i int) (bool, E) {
		return mapper(e, i), e
	})

}

func Take[E any](elements []E, n int) Iter[E] {
	var zero E

	iter := ToSliceIter(elements)

	return NewFuncIterator(iter, func(e E, _ int) (bool, E) {
		if n > 0 {
			n -= 1
			return true, e
		}

		return false, zero
	})
}
