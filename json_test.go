package benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func BenchmarkStdMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		if err := stdJsonMarshal(s); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkStdUnMar(b *testing.B) {
	var ss S
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		by, err := json.Marshal(s)
		if err != nil {
			b.Error(err)
			return
		}
		if err := stdJsonUnMarshal(by, ss); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkIterMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		if err := jsonIterMarshal(s); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkIterUnMar(b *testing.B) {
	var ss S
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		by, err := json.Marshal(s)
		if err != nil {
			b.Error(err)
			return
		}
		if err := jsonIterUnMarshal(by, ss); err != nil {
			b.Error(err)
			return
		}
	}
}
func BenchmarkMsgMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		if err := msgMarshal(s); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkMsgUnMar(b *testing.B) {
	var ss S
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		by, err := msgpack.Marshal(s)
		if err != nil {
			b.Error(err)
			return
		}
		if err := msgUnMarshal(by, ss); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkGobMar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		var net bytes.Buffer
		if err := gobMarshal(&net, s); err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkGobUnMar(b *testing.B) {
	var ss S
	for i := 0; i < b.N; i++ {
		s := S{
			Name: 1,
			Set:  2,
			User: "3",
		}
		var net bytes.Buffer
		if err := gob.NewEncoder(&net).Encode(s); err != nil {
			b.Error(err)
			return
		}
		if err := gobUnMarshal(&net, ss); err != nil {
			b.Error(err)
			return
		}
	}
}
