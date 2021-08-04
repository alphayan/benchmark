package benchmark

import (
	"testing"
)

func BenchmarkFgetjson(b *testing.B) {
	bb := []byte(testbyte)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !fjson(bb) {
			b.Error("fjson false")
		}
	}
}

func BenchmarkGgetjson(b *testing.B) {
	s := testbyte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !ggjson(s) {
			b.Error("gjson false")
		}
	}
}

func BenchmarkSgetjson(b *testing.B) {
	s := []byte(testbyte)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !sjson(s) {
			b.Error("sjson false")
		}
	}
}
