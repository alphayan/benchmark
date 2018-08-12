package benchmark

import (
	"bytes"
	"fmt"
	"strings"
)

func string1() string {
	s := ""
	for i := 0; i < 1000; i++ {
		s += "1"
	}
	return s
}

func string2() string {
	s := ""
	for i := 0; i < 1000; i++ {
		s = fmt.Sprintf("%s%s", s, "1")
	}
	return s
}
func string3() string {
	s := ""
	for i := 0; i < 1000; i++ {
		s = strings.Join([]string{s, "1"}, "")
	}
	return s
}

func string4() string {
	var buffer bytes.Buffer
	for i := 0; i < 1000; i++ {
		buffer.Write([]byte("1"))
	}
	return buffer.String()
}
