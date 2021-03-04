package benchmark

import (
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha11()
	}
}

func BenchmarkSHA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha12()
	}
}

func BenchmarkSHA3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha13()
	}
}

func BenchmarkSHA4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha14()
	}
}
