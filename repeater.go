package slicendice

import (
	"fmt"
	"log"
	"reflect"
)

// RepeatFunc invokes the given function `n` times
func RepeatFunc(n int, fn interface{}) error {
	fnVal := reflect.ValueOf(fn)
	if fnVal.Kind() != reflect.Func {
		return fmt.Errorf("provided argument is not a function")
	}

	for i := 0; i < n; i++ {
		fnVal.Call(nil)
	}

	return nil
}

// RepeatFuncWith invokes a given function `n` time
// with provided arguments and returns the result values
func RepeatFuncWith(n int, fn interface{}, args ...interface{}) ([]interface{}, error) {
	fnVal := reflect.ValueOf(fn)
	if fnVal.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided argument is not a function")
	}

	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	var returnValues []interface{}
	for i := 0; i < n; i++ {
		result := fnVal.Call(reflectArgs)
		if len(result) > 0 {
			returnValues = append(returnValues, result[0].Interface())
		}
	}

	return returnValues, nil
}

func Splice[E any](slice []E, start, deleteCount int, insert ...E) []E {
	if start < 0 {
		start = 0
	}

	if start > len(slice) {
		start = len(slice)
	}

	end := start + deleteCount
	if end > len(slice) {
		end = len(slice)
	}

	front := append([]E(nil), slice[:start]...)
	middle := append([]E(nil), insert...)
	back := append([]E(nil), slice[end:]...)

	// Combine them to form a new slice
	return append(append(front, middle...), back...)
}

func At[A any](slice []A, index int) A {
	n_slice := len(slice)

	if index == -1 && n_slice == 0 {
		index = 0
	}

	if index < 0 {
		index = n_slice + index
	}

	if index >= n_slice {
		log.Fatal("index out of bounds")
	}

	return slice[index]
}
