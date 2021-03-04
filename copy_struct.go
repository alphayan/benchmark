package benchmark

import (
	"reflect"

	"github.com/jinzhu/copier"
)

type St1 struct {
	ID   int
	Name string
}
type St2 struct {
	ID   int
	Name string
	Age  int
}

func copy1(st2 interface{}, st1 interface{}) {
	copier.Copy(st2, st1)
}
func copy2(st2 interface{}, st1 interface{}) {
	copyStruct(st2, st1)
}
func copyStruct(to interface{}, from interface{}) {
	ft := reflect.TypeOf(from)
	if ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
	}
	tv := reflect.ValueOf(to)
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}
	fv := reflect.ValueOf(from)
	if fv.Kind() == reflect.Ptr {
		fv = fv.Elem()
	}
	for i := 0; i < ft.NumField(); i++ {
		name := ft.Field(i).Name
		if fv.FieldByName(name).IsValid() {
			tv.FieldByName(name).Set(fv.FieldByName(name))
		}
	}
}
