package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Zip(t *testing.T) {
	t.Run("for empty array do nothing", func(t *testing.T) {
		ok, result := slicendice.Zip([]int{}, []int{})

		assert.Equal(t, true, ok)
		assert.Equal(t, [][]int{}, result)
	})

	t.Run("for arrays of differing lengths", func(t *testing.T) {
		ok, _ := slicendice.Zip([]int{1, 3}, []int{2})

		assert.Equal(t, false, ok)
	})

	t.Run("for array of same length", func(t *testing.T) {
		ok, results := slicendice.Zip([]int{1, 2}, []int{3, 4})

		assert.Equal(t, true, ok)
		assert.Equal(
			t,
			[][]int{
				{1, 3},
				{2, 4},
			},
			results,
		)
	})
}

func Test_ZipPairs(t *testing.T) {
	t.Run("for empty array do nothing", func(t *testing.T) {
		ok, result := slicendice.ZipPairs([]string{}, []int{})

		assert.Equal(t, true, ok)
		assert.Equal(t, 0, len(result))
		assert.Equal(t, []slicendice.Pair[string, int]{}, result)
	})

	t.Run("for arrays of differing lengths", func(t *testing.T) {
		ok, _ := slicendice.ZipPairs([]int{1, 3}, []int{2})

		assert.Equal(t, false, ok)
	})

	t.Run("for array of same length", func(t *testing.T) {
		ok, results := slicendice.ZipPairs([]string{"odd", "even"}, []int{3, 4})

		assert.Equal(t, true, ok)
		assert.Equal(
			t,
			[]slicendice.Pair[string, int]{
				{"odd", 3},
				{"even", 4},
			},
			results,
		)
	})
}
