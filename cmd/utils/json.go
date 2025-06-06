package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strings"
)

// MarshalStruct converts a struct into a JSON byte slice with sorted keys.
// It handles JSON struct tags including field renaming and omitempty directive.
//
// The function takes any struct as input and returns the JSON-encoded bytes and an error.
// Fields are sorted alphabetically by their JSON key names for consistent output.
//
// Struct tags are processed as follows:
// - `json:"-"` skips the field
// - `json:"name"` uses specified name as the JSON key
// - `json:",omitempty"` omits the field if it has zero/empty value
// - Unexported fields are always skipped
//
// Returns an error if:
// - Input is not a struct
// - Any field's value cannot be marshaled to JSON
//
// Example struct tags:
//
//	type Example struct {
//	  Field1 string `json:"f1,omitempty"`
//	  Field2 int    `json:"f2"`
//	  skip   bool   `json:"-"`
//	}
func MarshalStruct(v any) ([]byte, error) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// Check if the input is a struct
	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("MarshalStruct: expected a struct but got %s", rt.Kind())
	}

	type fieldEntry struct {
		Value any
		Key   string
	}
	var fields []fieldEntry

	for i := range rt.NumField() {
		field := rt.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		jsonKey := tagParts[0]
		if jsonKey == "" {
			jsonKey = field.Name
		}

		omitEmpty := slices.Contains(tagParts[1:], "omitempty")

		fv := rv.Field(i)

		if omitEmpty && isEmptyValue(fv) {
			continue
		}

		fields = append(fields, fieldEntry{
			Key:   jsonKey,
			Value: fv.Interface(),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Key < fields[j].Key
	})

	// Build JSON object manually
	buf := []byte{'{'}
	for i, f := range fields {
		valBytes, err := json.Marshal(f.Value)
		if err != nil {
			return nil, err
		}
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, fmt.Sprintf("%q:", f.Key)...)
		buf = append(buf, valBytes...)
	}
	buf = append(buf, '}')

	return buf, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.IsNil() || v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	default:
		return false
	}
}
