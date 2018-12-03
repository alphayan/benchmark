package benchmark

import "testing"

func BenchmarkUuidv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidv1()
	}
}
func BenchmarkUuidv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidv4()
	}
}
func BenchmarkUuid3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid3()
	}
}
