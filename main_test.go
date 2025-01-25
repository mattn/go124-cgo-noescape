package main

import (
	"testing"
)

func BenchmarkWithNoescape(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DoWrite1()
	}
}

func BenchmarkWithoutNoescape(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DoWrite2()
	}
}
