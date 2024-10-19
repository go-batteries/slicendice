package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Batch(t *testing.T) {
	t.Run("total < groups, returns [0, total - 1]", func(t *testing.T) {
		assert.Equal(t, [][]int{{0, 4}}, slicendice.Batch(5, 7))
	})

	t.Run("total == groups, returns [0, total - 1]", func(t *testing.T) {
		assert.Equal(
			t,
			[][]int{{0, 4}},
			slicendice.Batch(5, 5),
		)
	})

	t.Run("total and odd are not fully divisible", func(t *testing.T) {
		assert.Equal(
			t,
			[][]int{{0, 4}, {5, 9}, {10, 11}},
			slicendice.Batch(12, 5),
		)
	})

	t.Run("total and odd are divisble", func(t *testing.T) {
		assert.Equal(
			t,
			[][]int{{0, 2}, {3, 5}, {6, 8}, {9, 11}},
			slicendice.Batch(12, 3),
		)
	})
}

func Test_Repeat(t *testing.T) {
	t.Run("times == 0, return empty", func(t *testing.T) {
		repeated := slicendice.Repeat('?', 0)
		assert.Equal(t, 0, len(repeated))
	})

	t.Run("times == n, returns n items array", func(t *testing.T) {
		repeated := slicendice.Repeat("?", 5)
		assert.Equal(t, 5, len(repeated))
		assert.Equal(t, []string{"?", "?", "?", "?", "?"}, repeated)
	})
}
