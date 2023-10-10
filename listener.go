package fstore

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

type StoreListener struct {
	DontHash          []string
	Threshhold        int
	UseKeyCompression bool
	debug             bool
	database          Database
}

func Listener() *StoreListener {
	return &StoreListener{
		database: GetDatabase(),
		DontHash: []string{},
	}
}

func (s *StoreListener) log(l ...any) {
	if !s.debug {
		return
	}
	fmt.Print("[Listener] ")
	fmt.Println(l...)
}

func (s *StoreListener) EnableDebug() {
	s.debug = true
}

func (s *StoreListener) getStructValue(ref reflect.Value) any {
	fields := ref.NumField()
	result := map[string]any{}

	for i := 0; i < fields; i++ {
		field := ref.Field(i)
		jsonTag := ref.Type().Field(i).Tag.Get("json")
		name := strings.Replace(jsonTag, ",omitonempty", "", -1)
		name = strings.Replace(name, ",omitempty", "", -1)

		val := s.getFieldValue(field, name)
		if !isEmpty(val) {
			n := name
			if s.UseKeyCompression {
				n = s.database.SaveKey(name)
			}
			result[n] = val
		}
	}

	return result
}

func (s *StoreListener) getFieldValue(field reflect.Value, name string) any {
	if field.Kind() == reflect.String {
		str := field.String()
		if len(str) < s.Threshhold && !slices.Contains(s.DontHash, name) {
			return str
		}
		// s.log("=== hashing", name, "===")
		return s.database.SaveHash(str)
	}

	if field.Kind() == reflect.Struct {
		return s.getStructValue(field)
	}

	if field.Kind() == reflect.Slice {
		result := []any{}
		for i := 0; i < field.Len(); i++ {
			result = append(result, s.getFieldValue(field.Index(i), ""))
		}
		return result
	}

	if field.Kind() == reflect.Int {
		return field.Int()
	}

	s.log("== UNHANDLED", field.Kind())
	return field.Interface()
}

func (s *StoreListener) storeStruct(data interface{}) (any, error) {
	ref := reflect.ValueOf(data)

	result := s.getStructValue(ref)

	return result, nil
}

func (s *StoreListener) storeString(data string) (any, error) {
	s.log("Saving string", data)
	return data, nil
}
func (s *StoreListener) storeArray(data []any) (any, error) {
	s.log("Saving array")
	return data, nil
}

func (s *StoreListener) Store(data any) (any, error) {
	reflectVal := reflect.ValueOf(data)
	reflectKind := reflectVal.Kind()

	switch reflectKind {
	case reflect.Struct:
		return s.storeStruct(reflectVal.Interface())
	case reflect.Array:
		return s.storeArray(data.([]any))
	case reflect.Pointer:
		return s.storeStruct(reflectVal.Elem().Interface())
	case reflect.String:
		return s.storeString(data.(string))
	default:
		return nil, fmt.Errorf("could not store data of type %s", reflectKind.String())
	}
}
