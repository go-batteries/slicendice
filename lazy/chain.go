package lazy

// This is not in use
type ChainIter[E any] struct {
	iterators []Iter[E]
	current   int
}

func ChainIterators[E any](iterators ...Iter[E]) *ChainIter[E] {
	return &ChainIter[E]{iterators: iterators}
}

func (ci *ChainIter[E]) Next() (bool, E) {
	for ci.current < len(ci.iterators) {
		ok, value := ci.iterators[ci.current].Next()
		if ok {
			return true, value
		}

		ci.current++
	}

	var zero E
	return false, zero
}

func (ci *ChainIter[E]) Chain(iter Iter[E]) *ChainIter[E] {
	ci.iterators = append(ci.iterators, iter)
	return ci
}
