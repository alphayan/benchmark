package benchmark

import (
	"testing"
)

func BenchmarkMD51(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md51()
	}
}

func BenchmarkMD52(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md52()
	}
}
