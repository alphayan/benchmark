package benchmark

import (
	"fmt"
	"strconv"
)

func strconvitoa(i int) string {
	return strconv.Itoa(i)
}
func sprintfitoa(i int) string {
	return fmt.Sprintf("%d", i)
}
