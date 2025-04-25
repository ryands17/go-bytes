package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// PrintJSON pretty prints the given object as indented JSON to standard output.
// Marshaling errors are ignored.
func PrintJSON(obj any) {
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(bytes))
}

// CopyStructFields copies fields with matching names and types from the source struct to the destination struct.
// It uses reflection to iterate through all fields of the source struct and attempts to copy each one to
// a field with the same name in the destination struct.
//
// Parameters:
//   - dst: A pointer to the destination struct where fields will be copied to. Must be a pointer to a struct.
//   - src: A pointer to the source struct where fields will be copied from. Must be a pointer to a struct.
//
// The function only copies fields that:
//  1. Have the same name in both structs
//  2. Have the same type in both structs
//  3. Are settable in the destination struct
//
// Fields that don't meet these criteria are silently skipped.
//
// Example:
//
//	type Source struct {
//	    Name  string
//	    Age   int
//	    Email string
//	}
//
//	type Destination struct {
//	    Name  string
//	    Age   int
//	    Phone string  // Different field, won't be copied
//	}
//
//	src := Source{Name: "John", Age: 30, Email: "john@example.com"}
//	dst := Destination{}
//	CopyStructFields(&dst, &src)
//	// dst now contains {Name: "John", Age: 30, Phone: ""}
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
