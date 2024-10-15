package slicendice

func CombineK[E any](elements []E, k int) [][]E {
	n := len(elements)

	if n == 0 || k <= 0 || k > n {
		return [][]E{}
	}

	if n == 1 {
		return [][]E{{elements[0]}}
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
