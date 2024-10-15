package slicendice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factorial(t *testing.T) {
	t.Run("0! == 1", func(t *testing.T) {
		assert.Equal(t, int64(1), Factorial(0))
	})

	t.Run("1! == 1", func(t *testing.T) {
		assert.Equal(t, int64(1), Factorial(1))
	})

	t.Run("5! == 120", func(t *testing.T) {
		assert.Equal(t, int64(120), Factorial(5))
	})

	t.Run("11! == 11 * 10!", func(t *testing.T) {
		assert.Equal(t, 11*factorials[10], Factorial(11))
	})
}
