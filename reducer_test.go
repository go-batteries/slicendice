package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_ReduceArray(t *testing.T) {
	t.Run("behaves like a map, with empty accumulator", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		accumulator := []int{}
		multipler := func(acc []int, el int, _ int) []int {
			acc = append(acc, el*2)
			return acc
		}

		result := slicendice.Reduce(input, multipler, accumulator)

		assert.Equal(t, []int{2, 4, 6, 8}, result)
	})

	t.Run("behaves like a map, with non empty accumulator", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		accumulator := make([]int, len(input))

		multiplier := func(acc []int, el int, i int) []int {
			acc[i] = el * 2
			return acc
		}

		result := slicendice.Reduce(input, multiplier, accumulator)

		assert.Equal(t, []int{2, 4, 6, 8}, result)
	})

	t.Run("successfully reduces from one type to another type", func(t *testing.T) {
		input := []int{1, 2, 3, 4}

		adder := func(acc int, el int, _ int) int {
			return acc + el
		}

		result := slicendice.Reduce(input, adder, 0)

		assert.Equal(t, 1+2+3+4, result)
	})

	t.Run("behaves like a filter", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}

		onlyEvens := func(acc []int, el int, _ int) []int {
			if el%2 == 0 {
				acc = append(acc, el)
			}

			return acc
		}

		result := slicendice.Reduce(input, onlyEvens, []int{})

		assert.Equal(t, []int{2, 4, 6}, result)
	})
}
