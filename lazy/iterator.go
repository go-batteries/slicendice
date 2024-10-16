package lazy

type Iter[E any] interface {
	Next() (bool, E)
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

type Transform[E, V any] func(E, int) (bool, V)

type FuncIterator[E, V any] struct {
	iter        Iter[E]
	transformer Transform[E, V]
	index       int
}

func NewFuncIterator[E, V any](iter Iter[E], transformer Transform[E, V]) *FuncIterator[E, V] {
	return &FuncIterator[E, V]{
		iter:        iter,
		transformer: transformer,
		index:       0,
	}
}

func (fi *FuncIterator[E, V]) Next() (bool, V) {
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
