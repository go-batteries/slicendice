package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	t.Run("with empty array it return empty result", func(t *testing.T) {
		result := slicendice.Filter([]int{}, func(el int, _ int) bool {
			return false
		})

		assert.Equal(t, 0, len(result))
		assert.Equal(t, []int{}, result)
	})

	t.Run("filters even numbers from array", func(t *testing.T) {
		inputs := []int{1, 2, 3, 4, 5}

		result := slicendice.Filter(inputs, func(el int, _ int) bool {
			return el%2 == 0
		})

		assert.Equal(t, []int{2, 4}, result)
	})

}
