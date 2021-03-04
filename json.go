package benchmark

import (
	"encoding/gob"
	"encoding/json"
	"io"

	"github.com/vmihailenco/msgpack/v5"

	jsoniter "github.com/json-iterator/go"
)

type S struct {
	Name int
	Set  int
	User string
}

func stdJsonMarshal(s S) error {
	_, err := json.Marshal(s)
	return err
}

func jsonIterMarshal(s S) error {
	_, err := jsoniter.Marshal(s)
	return err
}
func stdJsonUnMarshal(b []byte, s S) error {
	return json.Unmarshal(b, &s)
}

func jsonIterUnMarshal(b []byte, s S) error {
	return jsoniter.Unmarshal(b, &s)
}

func msgMarshal(s S) error {
	_, err := msgpack.Marshal(s)
	return err
}
func msgUnMarshal(b []byte, s S) error {
	return msgpack.Unmarshal(b, &s)
}
func gobMarshal(w io.Writer, s S) error {
	return gob.NewEncoder(w).Encode(s)
}
func gobUnMarshal(r io.Reader, s S) error {
	return gob.NewDecoder(r).Decode(&s)
}
