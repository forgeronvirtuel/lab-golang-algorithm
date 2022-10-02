package _4_copy_arrays

import (
	"testing"
)

const length = 10000

func generateArray(length int) []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	return s
}

func BenchmarkConcat(b *testing.B) {
	s1 := generateArray(length)
	s2 := generateArray(length)
	for i := 0; i < b.N; i++ {
		ConcatWithCopy(s1, s2)
	}
}

func BenchmarkConcatWithFor(b *testing.B) {
	s1 := generateArray(length)
	s2 := generateArray(length)
	for i := 0; i < b.N; i++ {
		ConcatWithFor(s1, s2)
	}
}

func BenchmarkConcatWithAppend(b *testing.B) {
	s1 := generateArray(length)
	s2 := generateArray(length)
	for i := 0; i < b.N; i++ {
		ConcatWithAppend(s1, s2)
	}
}
