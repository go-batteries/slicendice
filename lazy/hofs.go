package lazy

import "github.com/go-batteries/slicendice"

func Map[E, V any](elements []E, mapper slicendice.MapperFunc[E, V]) Iter[V] {
	var zero V

	iter := ToSliceIter(elements)

	return NewFuncIterator(func() (bool, V) {
		ok, val := iter.Next()
		if !ok {
			return false, zero
		}

		result := mapper(val, iter.index)
		return true, result
	})
}

func Filter[E any](elements []E, filterer slicendice.FilterFunc[E]) Iter[E] {
	var zero E

	iter := ToSliceIter(elements)

	return NewFuncIterator(func() (bool, E) {
		for {
			ok, val := iter.Next()
			if !ok {
				return false, zero
			}

			if filterer(val, iter.index) {
				return true, val
			}
		}
	})
}

func Map2[E, V any](elements []E, mapper slicendice.MapperFunc[E, V]) Iter[V] {
	iter := ToSliceIter(elements)

	return NewFuncIterator2(iter, func(e E, _ int) (bool, V) {
		return true, mapper(e, iter.index)
	})

}

func Filter2[E any](elements []E, mapper slicendice.FilterFunc[E]) Iter[E] {
	iter := ToSliceIter(elements)

	return NewFuncIterator2(iter, func(e E, _ int) (bool, E) {
		return mapper(e, iter.index), e
	})

}
