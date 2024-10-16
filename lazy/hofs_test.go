package lazy_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-batteries/slicendice/lazy"
	"github.com/stretchr/testify/assert"
)

func Test_MapIntoIter(t *testing.T) {
	adder := func(el int, _ int) int {
		return el * 2
	}

	greeter := func(el string, _ int) string {
		return fmt.Sprintf("Hey %s!", el)
	}

	t.Run("empty array, next return false", func(t *testing.T) {
		mapper := lazy.MapIntoIter([]int{}, adder)

		ok, _ := mapper.Next()
		assert.False(t, ok)
	})

	t.Run("double item in array", func(t *testing.T) {
		input := []int{1, 3, 4}
		mapper := lazy.MapIntoIter(input, adder)
		counter := 0

		for ok, res := mapper.Next(); ok; ok, res = mapper.Next() {
			assert.Equal(t, adder(input[counter], 0), res)
			counter += 1
		}
	})

	t.Run("greets people in array", func(t *testing.T) {
		people := []string{"Dude", "Bro"}
		mapper := lazy.MapIntoIter(people, greeter)

		counter := 0

		for ok, res := mapper.Next(); ok; ok, res = mapper.Next() {
			assert.Equal(t, greeter(people[counter], 0), res)
			counter += 1
		}
	})
}

func Test_FilterIntoIter(t *testing.T) {
	t.Run("filter items one by one by one", func(t *testing.T) {
		input := []string{"a", "b", "a", "c"}
		afilters := func(el string, _ int) bool {
			return el == "a"
		}

		filterer := lazy.FilterIntoIter(input, afilters)
		res := []string{}

		for ok, out := filterer.Next(); ok; ok, out = filterer.Next() {
			res = append(res, out)
		}

		assert.Equal(t, []string{"a", "a"}, res)
	})
}

func Test_TakeIntoIter(t *testing.T) {
	t.Run("empty array takes none", func(t *testing.T) {
		input := []int{}

		iter := lazy.TakeIntoIter(input, 5)

		ok, next := iter.Next()

		assert.Equal(t, false, ok)
		assert.Empty(t, next)
	})

	t.Run("return all items when len(elements) < takeN", func(t *testing.T) {
		input := []int{1, 2}
		iter := lazy.TakeIntoIter(input, 5)

		result := []int{}

		for ok, e := iter.Next(); ok; ok, e = iter.Next() {
			result = append(result, e)
		}

		assert.Equal(t, len(input), len(result))
		assert.Equal(t, input, result)
	})

	t.Run("returns only takeN elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		iter := lazy.TakeIntoIter(input, 3)

		result := []int{}

		for ok, e := iter.Next(); ok; ok, e = iter.Next() {
			result = append(result, e)
		}

		assert.Equal(t, 3, len(result))
		assert.Equal(t, []int{1, 2, 3}, result)

	})
}

func Test_Map(t *testing.T) {
	inputs := []int{1, 2, 3, 4}

	inputIter := lazy.ToSliceIter(inputs)
	adder := func(el int, i int) string {
		return fmt.Sprintf("%d:%d", el, i)
	}

	iter := lazy.Map(inputIter, adder)

	count := 0

	for ok, res := iter.Next(); ok; ok, res = iter.Next() {
		splits := strings.Split(res, ":")

		assert.Equal(t, fmt.Sprintf("%d", inputs[count]), splits[0])
		assert.Equal(t, fmt.Sprintf("%d", count), splits[1])

		count += 1
	}
}

func Test_Filter(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	idxMap := map[int]int{1: 0, 2: 1, 3: 2, 4: 3}

	inputIter := lazy.ToSliceIter(inputs)
	evens := func(el int, i int) bool {
		assert.Equal(t, idxMap[el], i)
		return el%2 == 0
	}

	iter := lazy.Filter(inputIter, evens)
	result := []int{}

	for ok, next := iter.Next(); ok; ok, next = iter.Next() {
		result = append(result, next)
	}

	assert.Equal(t, []int{2, 4}, result)
}
