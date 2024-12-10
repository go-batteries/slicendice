package slicendice_test

import (
	"testing"

	"github.com/go-batteries/slicendice"
)

func Test_RepeatWithFunc(t *testing.T) {
	adder := func(num int, acc *int) {
		*acc += num
	}

	sum := 0

	slicendice.RepeatFunc(3, func() {
		adder(4, &sum)
	})

	if sum != 12 {
		t.Fatal("expected sum to be 12, got", sum)
	}
}

func Test_Splice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	newArr := slicendice.Splice(arr, 1, 1)

	if arr[1] == newArr[1] {
		t.Fatalf("arr and new array doesn't match. got %v, expected %v", arr, newArr)
	}
}
