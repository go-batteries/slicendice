package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Combine(t *testing.T) {
	items := []int{1, 2, 3}

	result := slicendice.Combine(items)

	assert.Equal(t, 7, len(result))

	expected := [][]int{
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}

	assert.Equal(t, expected, result)
}

func Test_CombineK(t *testing.T) {
	elements := []int{1, 2, 3, 4}
	k := 2
	combinations := slicendice.CombineK(elements, k)

	// for n = 4, taking upto k
	// 4C0 = 1 (blank)
	// 4C1 = 4
	// 4C2 = 6

	expected := [][]int{
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{4},
		{1, 4},
		{2, 4},
		{3, 4},
	}

	assert.Equal(t, 10, len(combinations))
	assert.Equal(t, expected, combinations)
}

func Test_Reverse(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		res := slicendice.Reverse([]int{})

		assert.Equal(t, 0, len(res))
	})

	t.Run("reverse the list", func(t *testing.T) {
		inputs := []int{5, 3, 4, 8, 1}

		outputs := slicendice.Reverse(inputs)

		assert.Equal(t, len(inputs), len(outputs))
		assert.Equal(t, []int{1, 8, 4, 3, 5}, outputs)
	})
}
