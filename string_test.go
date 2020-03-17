package benchmark

import (
	"bytes"
	"testing"
)

func BenchmarkString1(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		string1(s)
	}
}

func BenchmarkString2(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		string2(s)
	}
}
func BenchmarkString3(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		string3(s)
	}
}
func BenchmarkString4(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		string4(buf)
	}
}
