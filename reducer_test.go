package slicendice_test

import (
	"strings"
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

func Test_ReduceMap(t *testing.T) {
	t.Run("reduces an array to map", func(t *testing.T) {
		inputs := []string{"a1", "b2", "c3"}

		hashReduce := func(acc map[string]string, el string, _ int) map[string]string {
			splits := strings.Split(el, "")
			acc[splits[0]] = splits[1]
			return acc
		}

		results := slicendice.Reduce(inputs, hashReduce, map[string]string{})

		assert.Equal(t, map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		}, results)
	})
}

func Test_MapFilter(t *testing.T) {
	t.Run("filters odd func and multiples by 2", func(t *testing.T) {
		inputs := []int{1, 12, 3, 14, 5, 16}

		mapfilterFunc := func(v int, _ int) (int, bool) {
			if v%2 != 0 {
				return v * 2, true
			}

			return v, false
		}

		results := slicendice.MapFilter(inputs, mapfilterFunc)

		assert.Equal(t, []int{2, 6, 10}, results)
	})
}
