package fib

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(15)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(26)
	}
}

func BenchmarkFib25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(35)
	}
}
