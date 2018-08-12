package benchmark

import "testing"

func BenchmarkSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1()
	}
}
func BenchmarkSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice2()
	}
}
func BenchmarkSlice3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice3()
	}
}
