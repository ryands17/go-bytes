package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintJSON(obj any) {
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(bytes))
}

func CopyStructFields(dst, src any) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	for i := range srcVal.NumField() {
		srcField := srcVal.Field(i)
		srcFieldName := srcVal.Type().Field(i).Name

		dstField := dstVal.FieldByName(srcFieldName)
		if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcField.Type() {
			dstField.Set(srcField)
		}
	}
}
