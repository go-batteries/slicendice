package main

import (
	"fmt"
	"testing"
)

// Reverse by looping half the array
func reverseHalf(arr []int) {
	n := len(arr)
	// j := n - 1
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Reverse by looping the entire array
func reverseFull(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestReverseHalf(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	reverseHalf(arr)

	t.Log("Reverse Half")

	for i, val := range expected {
		if val != arr[i] {
			t.Errorf("values at index: %d didn't match %v\n", i, arr)
			t.Fail()
		}
	}
}

// Benchmark function
func BenchmarkReverseHalf(b *testing.B) {
	arr := make([]int, 10000)
	b.N = 190168

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseHalf(arr)
	}
}

func BenchmarkReverseFull(b *testing.B) {
	arr := make([]int, 10000)
	b.N = 190168

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseFull(arr)
	}
}

// Benchmark function
func BenchmarkReverseDynamicHalf(b *testing.B) {
	b.N = 190168
	sizes := []int{1000, 10000, 100000, 1000000}

	// Disable GC during benchmarking
	b.Setenv("GODEBUG", "gcstoptheworld=1")

	for _, size := range sizes {
		b.N = size

		b.Run(fmt.Sprintf("ArraySize=%d", size), func(b *testing.B) {
			arr := make([]int, b.N)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				reverseHalf(arr)
			}
		})
	}
}

func BenchmarkReverseDynamicFull(b *testing.B) {
	b.N = 190168

	sizes := []int{1000, 10000, 100000, 1000000}

	// Disable GC during benchmarking
	b.Setenv("GODEBUG", "gcstoptheworld=1")

	for _, size := range sizes {
		b.N = size

		b.Run(fmt.Sprintf("ArraySize=%d", size), func(b *testing.B) {
			arr := make([]int, b.N)

			b.ResetTimer() // Start timing after setup
			for i := 0; i < b.N; i++ {
				reverseFull(arr) // Or reverseFull(arr)
			}
		})
	}
}

// func BenchmarkReverseFull2(b *testing.B) {
// 	arr := make([]int, 10000)
// 	for i := 0; i < b.N; i++ {
// 		ReverseWithBlocking(arr, 256)
// 	}
// }
