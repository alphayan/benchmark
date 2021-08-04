package benchmark

import (
	"github.com/bitly/go-simplejson"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

var testbyte = `{"name":{"first":"Janet","last":"Prichard"},"age":"47"}`

func fjson(b []byte) bool {
	return fastjson.GetString(b, "age") == "47"
}
func ggjson(s string) bool {
	return gjson.Get(s, "age").String() == "47"
}
func sjson(b []byte) bool {
	res, _ := simplejson.NewJson(b)
	return res.Get("age").MustString() == "47"
}
