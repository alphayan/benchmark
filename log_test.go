package benchmark

import "testing"

func BenchmarkLogZap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logZap()
	}
}

func BenchmarkLogZapSugared(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logZapSugared()
	}
}

func BenchmarkLogZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logZerolog()
	}
}
