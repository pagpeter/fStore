package fstore

import (
	"reflect"
)

func isEmpty(s any) bool {
	switch v := s.(type) {
	case string:
		return v == ""
	case bool:
		return !v
	case int:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	default:
		rt := reflect.ValueOf(v)
		if reflect.Array == rt.Kind() || reflect.Slice == rt.Kind() || reflect.Map == rt.Kind() {
			return rt.Len() == 0
		}
		if reflect.Struct == rt.Kind() {
			return rt.NumField() == 0
		}

		return false
	}

}
