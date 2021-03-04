package benchmark

import "testing"

func BenchmarkAno(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ano()
	}
}

func BenchmarkUnAno(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unano()
	}
}
