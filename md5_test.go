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

func BenchmarkMD53(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md53()
	}
}

func BenchmarkMD54(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md54()
	}
}
