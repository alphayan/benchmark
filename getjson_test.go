package benchmark

import (
	"testing"
)

func BenchmarkFgetjson(b *testing.B) {
	bb := []byte(testbyte)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fjson(bb)
	}
}

func BenchmarkGgetjson(b *testing.B) {
	s := testbyte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ggjson(s)
	}
}
