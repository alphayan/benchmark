package benchmark

import (
	"testing"
)

func BenchmarkString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string1()
	}
}

func BenchmarkString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string2()
	}
}
func BenchmarkString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string3()
	}
}
func BenchmarkString4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string4()
	}
}
