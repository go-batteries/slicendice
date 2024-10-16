# slice and dice

A collection of commonly used higher order functions and data structures in go, like itertools.

- [x] Map
- [x] Filter
- [x] Reduce
- [x] Some
- [x] Every
- [x] Take
- [x] Permute
- [x] NextPermute
- [x] PermuteOrdered
- [x] Combine
- [x] Zip
- [x] ZipPairs

Lazy:

- [x] Iter
- [x] Map
- [x] Filter
- [ ] Reduce

Extra datastrcutures:

- [ ] Set

## Example

Although most of them are covered in `_test.go`, here is brief rundown.

```go

package main

import (
    "github.com/go-batteries/slicendice"
    "github.com/go-batteries/slicendice/lazy"
)

func main() {
    inputs := []int{1, 2, 3, 5}
    inputIter := lazy.ToSliceIter(inputs)

    adder := func(el int, _ int) int {
        return el + 1
    }

    evening := func(el int, _ int) bool {
        return el % 2 == 0
    }

    increments := slicendice.Map(inputs, adder)
    // returns []int{2,3,4,6}

    evens := slicendice.Filter(increments, evening)
    // return []int{2, 4, 6}

    iter := lazy.MapIntoIter(inputs, adder)
    // returns an iterator with Next()
    // or you could also do

    // iter := lazy.Map(inputIter, adder)

    // iterate
    for ok, res := iter.Next(); ok; ok, res = iter.Next() {
        // do something with result
    }

    // for a chained usecase
    // add one and filter evens
    // and take first 5
    inputs = []int{1,2,3,4,5,6,7,8,9}
    
    mapper := lazy.Map(lazy.ToSliceIter(inputs), adder)
    filterr := lazy.Filter(mapper, evening)

    iter2 := lazy.Take(5, filterr)
    // or
    iter2 = lazy.Take(5, lazy.Filter(lazy.Map(inputIter, added), evening))
}

```
