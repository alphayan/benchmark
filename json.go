package benchmark

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var str = `{"name":1,"sex":2,"user":3}`

func stdJsonMarshal() {
	json.Marshal(str)
}

func jsonIterMarshal() {
	jsoniter.Marshal(str)
}
func stdJsonUnMarshal() {
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
}

func jsonIterUnMarshal() {
	var m map[string]interface{}
	jsoniter.Unmarshal([]byte(str), &m)
}
