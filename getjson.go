package benchmark

import (
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

var testbyte = `{"name":{"first":"Janet","last":"Prichard"},"age":"47"}`

func fjson(b []byte) string {
	return fastjson.GetString(b, "age")
}
func ggjson(s string) string {
	return gjson.Get(s, "age").String()
}
