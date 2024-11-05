package main

import (
	"testing"
)

// Reverse by looping half the array
func reverseHalf(arr []int) {
	n := len(arr)
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

func ReverseWithBlocking(arr []int, blockSize int) {
	n := len(arr)
	for i := 0; i < n; i += blockSize {
		end := i + blockSize
		if end > n {
			end = n
		}
		// Reverse elements within the block
		reverseBlock(arr, i, end-1)
	}
}

func reverseBlock(arr []int, start, end int) {
	for i := start; i < (start+end)/2; i++ {
		j := end - (i - start) - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	//
	// for start < end {
	// 	arr[start], arr[end] = arr[end], arr[start]
	// 	start++
	// 	end--
	// }
}

// Benchmark function
func BenchmarkReverseHalf(b *testing.B) {
	arr := make([]int, 10000) // Adjust size for testing
	for i := 0; i < b.N; i++ {
		reverseHalf(arr)
	}
}

func BenchmarkReverseFull(b *testing.B) {
	arr := make([]int, 10000) // Adjust size for testing
	for i := 0; i < b.N; i++ {
		reverseFull(arr)
	}
}

func BenchmarkReverseFull2(b *testing.B) {
	arr := make([]int, 10000) // Adjust size for testing
	for i := 0; i < b.N; i++ {
		ReverseWithBlocking(arr, 512)
	}
}
