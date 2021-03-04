package benchmark

import "math/rand"

var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = longLetters[rand.Intn(len(longLetters))]
	}
	return b
}

func randUp(n int) []byte {
	if n <= 0 {
		return nil
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return nil
	}
	for i, x := range b {
		arc = x & 61
		b[i] = longLetters[arc]
	}
	return b
}
