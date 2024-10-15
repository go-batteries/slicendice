package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
	"github.com/stretchr/testify/assert"
)

func Test_Permutations(t *testing.T) {
	items := []int{1, 2, 4}
	expected := [][]int{
		{1, 2, 4},
		{1, 4, 2},
		{2, 1, 4},
		{2, 4, 1},
		{4, 1, 2},
		{4, 2, 1},
	}

	ok, results := slicendice.Permutations(items)

	assert.Equal(t, true, ok)
	assert.Equal(t, expected, results)
}

func Test_PermutationsOrdered(t *testing.T) {
	items := []int{1, 2, 4}
	expected := [][]int{
		{1, 2, 4},
		{1, 4, 2},
		{2, 1, 4},
		{2, 4, 1},
		{4, 1, 2},
		{4, 2, 1},
	}

	ok, results := slicendice.PermutationsOrdered(items)

	assert.Equal(t, true, ok)
	assert.Equal(t, expected, results)
}

func Test_NextPermute(t *testing.T) {
	items := []int{1, 2, 4}

	_, next := slicendice.NextPermute(items)
	assert.Equal(t, []int{1, 4, 2}, next)

	_, next = slicendice.NextPermute(next)
	assert.Equal(t, []int{2, 1, 4}, next)

	// for i := 0; i < len(results)-1; i++ {
	// 	ok, next := slicendice.NextPermute(results[i])
	// 	assert.Equal(t, true, ok)
	//
	// 	assert.Equal(t, results[i+1], next, fmt.Sprintf("failed at %d", i))
	// }

}
