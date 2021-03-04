package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRand1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomString(10)
	}
}
func BenchmarkRand2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randUp(10)
	}
}
