package benchmark

import (
	"crypto/md5"
	"fmt"
)

func md51() string {
	return fmt.Sprintf("%x", md5.Sum([]byte("abc123")))
}

func md52() string {
	w := md5.New()
	w.Write([]byte("abc123"))
	return fmt.Sprintf("%x", w.Sum(nil))
}
