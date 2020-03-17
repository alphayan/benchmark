package benchmark

import "testing"

func BenchmarkChannel1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channel1()
	}

}

func BenchmarkChanSync2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sync2()
	}
}
