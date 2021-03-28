package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkItoaStrconv(b *testing.B) {
	ii := 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(ii)
	}
}
func BenchmarkItoaSprintf(b *testing.B) {
	ii := 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", ii)
	}
}
