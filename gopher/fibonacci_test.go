package main

import (
	"testing"
)

func TestFibonacciAt(t *testing.T) {
	tt := []struct {
		f uint16
		e uint64
	}{
		{f: 0, e: 0},
		{f: 1, e: 1},
		{f: 2, e: 1},
		{f: 3, e: 2},
		{f: 4, e: 3},
		{f: 8, e: 21},
		{f: 16, e: 987},
		{f: 32, e: 15621},
	}
	for _, tc := range tt {
		r := fibonacciAt(tc.f)
		if r != tc.e {
			t.Errorf("expected '%d' for fib(%d), got '%d'", tc.e, tc.f, r)
		}
	}
}

func BenchmarkFibonacciAt2(b *testing.B) {
	benchmarkFibonacciAt(b, 2)
}

func BenchmarkFibonacciAt4(b *testing.B) {
	benchmarkFibonacciAt(b, 4)
}

func BenchmarkFibonacciAt8(b *testing.B) {
	benchmarkFibonacciAt(b, 8)
}

func BenchmarkFibonacciAt16(b *testing.B) {
	benchmarkFibonacciAt(b, 16)
}

func BenchmarkFibonacciAt32(b *testing.B) {
	benchmarkFibonacciAt(b, 32)
}

func benchmarkFibonacciAt(b *testing.B, n uint16) {
	for i := 0; i < b.N; i++ {
		fibonacciAt(n)
	}
}
