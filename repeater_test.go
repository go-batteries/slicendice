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
