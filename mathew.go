package slicendice

import "cmp"

func Max[E cmp.Ordered](values ...E) (e E) {
	if len(values) == 0 {
		return e
	}

	if len(values) == 1 {
		return values[0]
	}

	newMax := values[0]
	n := len(values)

	for i := 0; i < n/2; i++ {
		j := n - i - 1

		el := values[i]
		ol := values[j]

		if el > newMax {
			newMax = el
		}
		if ol > newMax {
			newMax = ol
		}
	}

	if n%2 != 0 {
		secondMid := values[n/2+1]
		if secondMid > newMax {
			return secondMid
		}
	}

	return newMax
}

func Min[E cmp.Ordered](values ...E) (e E) {
	if len(values) == 0 {
		return e
	}

	if len(values) == 1 {
		return values[0]
	}

	newMin := values[0]
	n := len(values)

	for i := 0; i < n/2; i++ {
		j := n - i - 1

		el := values[i]
		ol := values[j]

		if el < newMin {
			newMin = el
		}
		if ol < newMin {
			newMin = ol
		}
	}

	if n%2 != 0 {
		secondMid := values[n/2+1]
		if secondMid < newMin {
			return secondMid
		}
	}

	return newMin
}
