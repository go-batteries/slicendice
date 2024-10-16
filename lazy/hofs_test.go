package lazy_test

import (
	"fmt"
	"testing"

	"github.com/go-batteries/slicendice/lazy"
	"github.com/stretchr/testify/assert"
)

func Test_MapIter(t *testing.T) {
	adder := func(el int, _ int) int {
		return el * 2
	}

	greeter := func(el string, _ int) string {
		return fmt.Sprintf("Hey %s!", el)
	}

	t.Run("empty array, next return false", func(t *testing.T) {
		mapper := lazy.Map([]int{}, adder)

		ok, _ := mapper.Next()
		assert.False(t, ok)
	})

	t.Run("double item in array", func(t *testing.T) {
		input := []int{1, 3, 4}
		mapper := lazy.Map(input, adder)
		counter := 0

		for ok, res := mapper.Next(); ok; ok, res = mapper.Next() {
			assert.Equal(t, adder(input[counter], 0), res)
			counter += 1
		}
	})

	t.Run("double item in array", func(t *testing.T) {
		input := []int{1, 3, 4}
		mapper := lazy.Map2(input, adder)
		counter := 0

		for ok, res := mapper.Next(); ok; ok, res = mapper.Next() {
			assert.Equal(t, adder(input[counter], 0), res)
			counter += 1
		}
	})

	t.Run("greets people in array", func(t *testing.T) {
		people := []string{"Dude", "Bro"}
		mapper := lazy.Map(people, greeter)

		counter := 0

		for ok, res := mapper.Next(); ok; ok, res = mapper.Next() {
			assert.Equal(t, greeter(people[counter], 0), res)
			counter += 1
		}
	})
}

func Test_FilterIter(t *testing.T) {
	t.Run("filter items one by one by one", func(t *testing.T) {
		input := []string{"a", "b", "a", "c"}
		afilters := func(el string, _ int) bool {
			return el == "a"
		}

		filterer := lazy.Filter(input, afilters)
		res := []string{}

		for ok, out := filterer.Next(); ok; ok, out = filterer.Next() {
			res = append(res, out)
		}

		assert.Equal(t, []string{"a", "a"}, res)
	})

	t.Run("filter items one by one by one", func(t *testing.T) {
		input := []string{"a", "b", "a", "c"}
		afilters := func(el string, _ int) bool {
			return el == "a"
		}

		filterer := lazy.Filter2(input, afilters)
		res := []string{}

		for ok, out := filterer.Next(); ok; ok, out = filterer.Next() {
			res = append(res, out)
		}

		assert.Equal(t, []string{"a", "a"}, res)
	})

}
