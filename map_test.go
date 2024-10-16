package slicendice_test

import (
	"fmt"
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	adder := func(el int, _ int) int {
		return el * 2
	}

	t.Run("with empty array it returns empty", func(t *testing.T) {
		result := slicendice.Map([]int{}, adder)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, []int{}, result)
	})

	t.Run("with non empty slice, returned array is of same size", func(t *testing.T) {
		input := []int{1, 2, 3, 4}

		result := slicendice.Map(input, adder)

		assert.Equal(t, 4, len(result))
		assert.Equal(t, []int{2, 4, 6, 8}, result)
	})

	t.Run("map from one type to another", func(t *testing.T) {
		input := []int{1, 2, 3}
		convert := func(el int, _ int) string {
			return fmt.Sprintf("%d", el)
		}

		result := slicendice.Map(input, convert)
		assert.Equal(t, []string{"1", "2", "3"}, result)
	})
}
