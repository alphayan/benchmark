package benchmark

import "testing"

//func TestUuid3(t *testing.T) {
//	t.Log(uuidv1())
//	t.Log(uuidv4())
//	t.Log(uuid3())
//	t.Log(guuid())
//}

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
func BenchmarkGUuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guuid()
	}
}
