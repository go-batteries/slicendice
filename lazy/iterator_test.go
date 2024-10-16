package lazy_test

import (
	"testing"

	"github.com/go-batteries/slicendice/lazy"
	"github.com/stretchr/testify/assert"
)

func Test_Iterator(t *testing.T) {
	array := []int{1, 2, 3, 4}
	it := lazy.ToSliceIter(array)

	t.Run("using loops", func(t *testing.T) {
		for i := 0; i < len(array); i++ {
			ok, res := it.Next()

			assert.Equal(t, true, ok)
			assert.Equal(t, array[i], res)
		}

		ok, _ := it.Next()
		assert.Equal(t, false, ok)
	})

	t.Run("using iterator", func(t *testing.T) {
		it.Reset()

		count := 0

		for ok, el := it.Next(); ok; ok, el = it.Next() {
			assert.Equal(t, array[count], el)
			count += 1
		}
	})
}
