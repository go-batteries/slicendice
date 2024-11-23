package slicendice

import (
	"cmp"
)

func CombineK[E any](elements []E, k int) [][]E {
	n := len(elements)

	if n == 0 || k <= 0 || k > n {
		return [][]E{}
	}

	if n == 1 {
		return [][]E{elements}
	}

	totalCombines := 1 << n
	// result := make([][]E, totalCombines)
	result := [][]E{}

	for mask := 0; mask < totalCombines; mask++ {
		combines := []E{}
		count := 0

		for i, el := range elements {
			if mask&(1<<i) != 0 {
				combines = append(combines, el)
				count++
			}
		}

		if count <= k {
			result = append(result, combines)
		}
	}

	// remove the first blank element
	return result[1:]
}

func Combine[E any](elements []E) [][]E {
	return CombineK(elements, len(elements))
}

type permuteStackEntry[E cmp.Ordered] struct {
	mask      int
	currState []E
}

type stack[E cmp.Ordered] []*permuteStackEntry[E]

func (s *stack[E]) Push(v *permuteStackEntry[E]) {
	*s = append(*s, v)
}

func (s *stack[E]) Pop() (*permuteStackEntry[E], bool) {
	old := *s
	l := len(old)

	if l == 0 {
		return nil, false
	}

	item := old[l-1]
	*s = old[0 : l-1]

	old[l-1] = nil

	return item, true
}

// Permutations gets all the possible permutations
// This version doesn't gurrantee lexical ordered
func Permutations[E cmp.Ordered](elements []E) (bool, [][]E) {
	n := len(elements)

	if n == 0 {
		return true, [][]E{elements}
	}

	result := [][]E{}
	bucket := stack[E]{}
	bucket.Push(&permuteStackEntry[E]{0, []E{}})

	for len(bucket) > 0 {
		item, ok := bucket.Pop()
		if !ok { // shouldn't have reached here
			return false, [][]E{elements}
		}

		if len(item.currState) == n {
			result = append(result, item.currState)
			continue
		}

		for i := n - 1; i >= 0; i-- {
			el := elements[i]
			if item.mask&(1<<i) == 0 {
				bucket.Push(&permuteStackEntry[E]{
					mask:      item.mask | (1 << i),
					currState: append(item.currState, el),
				})
			}
		}
	}

	return true, result
}

const MAX_COUNT = 10000000

// PermutationsOrdered same as Permutations but lexically ordered
func PermutationsOrdered[E cmp.Ordered](elements []E) (bool, [][]E) {
	n := len(elements)
	if n == 0 {
		return true, [][]E{elements}
	}

	current := make([]E, n)
	copy(current, elements)

	result := [][]E{}
	result = append(result, current)

	for totalPerms := Factorial(n); totalPerms >= 0; totalPerms-- {
		ok, next := NextPermute(current)
		if !ok {
			// No more permutations
			break
		}

		result = append(result, next)
		current = next
	}

	return true, result
}

// NextPermute find the next permuted if possible
// Returns ok, elements. If next permutation is exhausted
// ok is false
// The algorithm goes like:
// - Find the largest index of k such that a[k] < a[k+1]
// - Find the largest index l such that a[l] > a[k]
// - Swap the value of a[l] and a[k], and
// - Reverse the list
func NextPermute[E cmp.Ordered](elements []E) (bool, []E) {
	n := len(elements)
	if n == 0 {
		return true, []E{}
	}

	copied := make([]E, n)

	copy(copied, elements)

	// Step 1
	k := n - 2
	for k >= 0 && copied[k] >= copied[k+1] {
		k -= 1
	}

	if k < 0 {
		return false, copied
	}

	// Step 2
	l := n - 1
	for l >= 0 && copied[k] >= copied[l] {
		l -= 1
	}

	copied[k], copied[l] = copied[l], copied[k]

	// Step 3
	reversedPart := Reverse(copied[k+1:])
	copied = append(copied[:k+1], reversedPart...)

	return true, copied
}

type Comparator[E cmp.Ordered] func(prev, next E) bool

// NextPermuteFunc, works like NextPermute
// but allows you to provide a comparator func
// default should be `>=`
func NextPermuteFunc[E cmp.Ordered](elements []E, comparator Comparator[E]) (bool, []E) {
	n := len(elements)
	if n == 0 {
		return true, []E{}
	}

	copied := make([]E, n)

	copy(copied, elements)

	// Step 1
	k := n - 2
	for k >= 0 && comparator(copied[k], copied[k+1]) {
		k -= 1
	}

	if k < 0 {
		return false, copied
	}

	// Step 2
	l := n - 1
	for l >= 0 && comparator(copied[k], copied[l]) {
		l -= 1
	}

	copied[k], copied[l] = copied[l], copied[k]

	// Step 3
	reversedPart := Reverse(copied[k+1:])
	copied = append(copied[:k+1], reversedPart...)

	return true, copied
}

func Reverse[E any](elements []E) []E {
	n := len(elements)

	if n == 0 {
		return []E{}
	}

	out := make([]E, len(elements))

	copy(out, elements)

	half := n / 2

	for i := 0; i < half; i++ {
		j := n - i - 1
		out[i], out[j] = out[j], out[i]
	}

	return out
}
