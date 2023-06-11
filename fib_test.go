package fib

import (
	"testing"
)

type Alloc struct {
	N *int
}

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := Alloc{
			N: &i,
		}
		_ = Fib(20 + uint(*a.N*0))
	}
}

func BenchmarkFib25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(25)
	}
}
