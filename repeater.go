package slicendice

import (
	"fmt"
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
