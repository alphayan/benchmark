package benchmark

import "testing"

func BenchmarkCopy1(b *testing.B) {
	var s1 = &St1{
		ID:   1,
		Name: "test",
	}
	var s2 = &St2{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy1(s2, s1)
	}
}

func BenchmarkCopy2(b *testing.B) {
	var s1 = &St1{
		ID:   1,
		Name: "test",
	}
	var s2 = &St2{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy2(s2, s1)
	}
}
