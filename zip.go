package slicendice

// Zips two arrays of equal length and same type
// into a single array, such that
// f(left, right) = [left[i], right[i]]
func Zip[E any](left, right []E) (bool, [][]E) {
	n_l, n_r := len(left), len(right)

	if n_l != n_r {
		return false, [][]E{}
	}

	if n_l == 0 {
		return true, [][]E{}
	}

	res := make([][]E, n_l)

	for i, el := range left {
		res[i] = []E{el, right[i]}
	}

	return true, res
}

type Pair[E, V any] struct {
	First E
	Last  V
}

// ZipPairs is same as Zip but works on
// heterogeneous types liek []string{}, []int{}
func ZipPairs[E, V any](left []E, right []V) (bool, []Pair[E, V]) {
	n_l, n_r := len(left), len(right)

	if n_l != n_r {
		return false, []Pair[E, V]{}
	}

	if n_l == 0 {
		return true, []Pair[E, V]{}
	}

	res := make([]Pair[E, V], n_l)

	for i, el := range left {
		res[i] = Pair[E, V]{
			First: el,
			Last:  right[i],
		}
	}

	return true, res
}
