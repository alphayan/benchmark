package benchmark

import "testing"

func BenchmarkStdMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdJsonMarshal()
	}
}

func BenchmarkStdUnMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdJsonUnMarshal()
	}
}

func BenchmarkIterMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonIterMarshal()
	}
}

func BenchmarkIterUnMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonIterUnMarshal()
	}
}
