package slicendice

type Set[V comparable] map[V]struct{}

// ToSet creates a new set from given values
func ToSet[V comparable](values ...V) Set[V] {
	s := Set[V]{}
	for _, value := range values {
		s[value] = struct{}{}
	}
	return s
}

// Add adds a value to the set
func (s Set[V]) Add(value V) {
	s[value] = struct{}{}
}

// Remove removes a value from the set
func (s Set[V]) Remove(value V) {
	delete(s, value)
}

// Has checks if the set contains a value
func (s Set[V]) Has(value V) bool {
	_, ok := s[value]
	return ok
}

// ToList returns the elements as a slice
func (s Set[V]) ToList() []V {
	out := make([]V, 0, len(s))
	for v := range s {
		out = append(out, v)
	}
	return out
}

// Diff returns elements in s but not in other
func (s Set[V]) Diff(other Set[V]) Set[V] {
	diff := Set[V]{}
	for v := range s {
		if _, exists := other[v]; !exists {
			diff[v] = struct{}{}
		}
	}
	return diff
}

type DifferFunc[V comparable] func(a, b V) bool

// DiffWith returns elements in s where the differ func returns true
func (s Set[V]) DiffWith(other Set[V], differ DifferFunc[V]) Set[V] {
	diff := Set[V]{}
	for v := range s {
		if o, exists := findMatching(other, v, differ); !exists {
			diff[v] = struct{}{}
		} else {
			_ = o
		}
	}
	return diff
}

func findMatching[V comparable](s Set[V], val V, eq func(a, b V) bool) (V, bool) {
	for otherVal := range s {
		if eq(val, otherVal) {
			return otherVal, true
		}
	}
	var zero V
	return zero, false
}

// Intersection returns a new set with elements in both sets
func (s Set[V]) Intersection(other Set[V]) Set[V] {
	result := Set[V]{}
	for v := range s {
		if _, ok := other[v]; ok {
			result[v] = struct{}{}
		}
	}
	return result
}

// Union returns a new set with all elements from both sets
func (s Set[V]) Union(other Set[V]) Set[V] {
	union := Set[V]{}
	for v := range s {
		union[v] = struct{}{}
	}
	for v := range other {
		union[v] = struct{}{}
	}
	return union
}
