package lazy

type Iter[E any] interface {
	Next() (bool, E)
}

type ChainedIter[E, V any] interface {
	Iter[E]
	Chain(Iter[E]) Iter[V]
}

type SliceIterator[E any] struct {
	elements []E

	index int
	n     int // length
}

func (it *SliceIterator[E]) Next() (bool, E) {
	if it.index >= it.n {
		var zero E
		return false, zero
	}

	el := it.elements[it.index]
	it.index += 1

	return true, el
}

func (it *SliceIterator[E]) Reset() {
	it.index = 0
}

type FuncIterator[E any] struct {
	next func() (bool, E)
}

func NewFuncIterator[E any](nextFunc func() (bool, E)) *FuncIterator[E] {
	return &FuncIterator[E]{next: nextFunc}
}

func (fiter *FuncIterator[E]) Next() (bool, E) {
	return fiter.next()
}

type Transform[E, V any] func(E, int) (bool, V)

type FuncIterator2[E, V any] struct {
	iter        Iter[E]
	transformer Transform[E, V]
	index       int
}

func NewFuncIterator2[E, V any](iter Iter[E], transformer Transform[E, V]) *FuncIterator2[E, V] {
	return &FuncIterator2[E, V]{
		iter:        iter,
		transformer: transformer,
		index:       0,
	}
}

func (fi *FuncIterator2[E, V]) Next() (bool, V) {
	for {
		ok, value := fi.iter.Next()
		if !ok {
			var zero V
			return false, zero // No more elements to iterate
		}

		valid, newValue := fi.transformer(value, fi.index)
		fi.index++

		if valid {
			return true, newValue
		}
	}
}

func ToSliceIter[E any](elements []E) *SliceIterator[E] {
	return &SliceIterator[E]{
		elements: elements,
		index:    0,
		n:        len(elements),
	}
}
