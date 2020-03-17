package benchmark

import (
	"bytes"
	"fmt"
	"strings"
)

func string1(s string) string {
	for i := 0; i < 1000; i++ {
		s += "1"
	}
	return s
}

func string2(s string) string {
	for i := 0; i < 1000; i++ {
		s = fmt.Sprintf("%s%s", s, "1")
	}
	return s
}
func string3(s string) string {
	for i := 0; i < 1000; i++ {
		s = strings.Join([]string{s, "1"}, "")
	}
	return s
}

func string4(buffer bytes.Buffer) string {
	for i := 0; i < 1000; i++ {
		buffer.Write([]byte("1"))
	}
	return buffer.String()
}
