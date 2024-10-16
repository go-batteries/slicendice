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

type FuncIterator[E any] struct {
	next func() (bool, E)
}

func NewFuncIterator[E any](nextFunc func() (bool, E)) *FuncIterator[E] {
	return &FuncIterator[E]{next: nextFunc}
}

func (fiter *FuncIterator[E]) Next() (bool, E) {
	return fiter.next()
}

func ToSliceIter[E any](elements []E) *SliceIterator[E] {
	return &SliceIterator[E]{
		elements: elements,
		index:    0,
		n:        len(elements),
	}
}
