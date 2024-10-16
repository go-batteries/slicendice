package lazy

import "github.com/go-batteries/slicendice"

func MapIntoIter[E, V any](elements []E, mapper slicendice.MapperFunc[E, V]) Iter[V] {
	iter := ToSliceIter(elements)

	return NewFuncIterator(iter, func(e E, _ int) (bool, V) {
		return true, mapper(e, iter.index)
	})

}

func Map[E, V any](iter Iter[E], mapper slicendice.MapperFunc[E, V]) Iter[V] {
	return NewFuncIterator(iter, func(e E, i int) (bool, V) {
		return true, mapper(e, i)
	})

}

func FilterIntoIter[E any](elements []E, mapper slicendice.FilterFunc[E]) Iter[E] {
	iter := ToSliceIter(elements)

	return NewFuncIterator(iter, func(e E, _ int) (bool, E) {
		return mapper(e, iter.index), e
	})

}

func Filter[E any](iter Iter[E], mapper slicendice.FilterFunc[E]) Iter[E] {
	return NewFuncIterator(iter, func(e E, i int) (bool, E) {
		return mapper(e, i), e
	})

}

func TakeIntoIter[E any](n int, elements []E) Iter[E] {
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

func Take[E any](n int, iter Iter[E]) Iter[E] {
	var zero E

	return NewFuncIterator(iter, func(e E, _ int) (bool, E) {
		if n > 0 {
			n -= 1
			return true, e
		}

		return false, zero
	})

}
