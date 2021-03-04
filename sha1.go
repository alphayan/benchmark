package benchmark

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func sha11() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte("abc123")))
}

func sha12() string {
	w := sha1.New()
	w.Write([]byte("abc123"))
	return fmt.Sprintf("%x", w.Sum(nil))
}

func sha13() string {
	b := sha1.Sum([]byte("abc123"))
	return hex.EncodeToString(b[:])
}
func sha14() string {
	w := sha1.New()
	w.Write([]byte("abc123"))
	return hex.EncodeToString(w.Sum(nil))
}
