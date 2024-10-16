package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChainMap(t *testing.T) {
	t.Run("add and double", func(t *testing.T) {
		inputs := ToSliceIter([]int{1, 2, 3, 4, 5, 6, 7, 8})

		adder := func(el int, _ int) int {
			return el + 1
		}

		mutter := func(el int, _ int) int {
			return el * 2
		}

		takeN := 3

		it := Take(
			takeN,
			Map(Map(inputs, adder), mutter),
		)

		result := []int{}

		for ok, res := it.Next(); ok; ok, res = it.Next() {
			result = append(result, res)
		}

		assert.Equal(t, takeN, len(result))
		assert.Equal(t, []int{4, 6, 8}, result)

		ok, _ := it.Next()
		assert.Equal(t, false, ok)

	})

	t.Run("add one and filter even", func(t *testing.T) {
		inputs := ToSliceIter([]int{21, 13, 12, 15, 51, 20})

		adder := func(el int, _ int) int {
			return el + 1
		}

		filtr := func(el int, _ int) bool {
			return el%2 == 0
		}

		it := Take(3, Filter(Map(inputs, adder), filtr))

		results := []int{}

		for ok, res := it.Next(); ok; ok, res = it.Next() {
			results = append(results, res)
		}

		assert.Equal(t, 3, len(results))
		assert.Equal(t, []int{22, 14, 16}, results)

		ok, _ := it.Next()
		assert.Equal(t, false, ok)

	})

}
