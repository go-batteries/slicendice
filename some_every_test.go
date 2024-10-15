package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Some(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}

	t.Run("for an empty array, return false", func(t *testing.T) {
		result := slicendice.Some([]int{}, func(el int, index int) bool {
			return el%2 == 0
		})

		assert.Equal(t, false, result)
	})

	t.Run("should return true if one match", func(t *testing.T) {

		hasEven := slicendice.Some(ints, func(el int, index int) bool {
			return el%2 == 0
		})

		assert.Equal(t, true, hasEven)
	})

	t.Run("should return false if none match", func(t *testing.T) {
		hasNegative := slicendice.Some(ints, func(el int, index int) bool {
			return el < 0
		})

		assert.Equal(t, false, hasNegative)
	})

	t.Run("returns false if none match", func(t *testing.T) {
		hasOdd := slicendice.Some([]int{2, 4}, func(el int, _ int) bool {
			return el%2 != 0
		})

		assert.Equal(t, false, hasOdd)
	})

}

func TestEvery(t *testing.T) {
	t.Run("returns false if one condition fails", func(t *testing.T) {
		ints := []int{1, -2, 3}

		allPositive := slicendice.Every(ints, func(el int, index int) bool {
			return el > 0
		})

		assert.Equal(t, false, allPositive)
	})

	t.Run("returns true if all conditions are meet", func(t *testing.T) {
		allEven := slicendice.Every([]int{2, 4, 6}, func(el int, index int) bool {
			return el%2 == 0
		})

		assert.Equal(t, true, allEven)
	})

}
