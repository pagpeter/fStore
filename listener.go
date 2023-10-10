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

func (s *StoreListener) getMapValue(ref reflect.Value) any {
	fields := ref.MapKeys()

	result := map[string]any{}
	for _, fieldName := range fields {
		field := ref.MapIndex(fieldName)
		s.log("tmp:", field, field.Elem(), field.Type())
		name := fieldName.String()
		s.log(name)
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
	} else if field.Kind() == reflect.Float64 {
		i := field.Float()
		if !slices.Contains(s.DontHash, name) {
			return i
		}
		return s.database.SaveHash(fmt.Sprintf("%v", i))
	} else if field.Kind() == reflect.Int64 {
		i := field.Int()
		if !slices.Contains(s.DontHash, name) {
			return i
		}
		return s.database.SaveHash(fmt.Sprintf("%v", i))
	} else if field.Kind() == reflect.Struct {
		return s.getStructValue(field)
	} else if field.Kind() == reflect.Slice {
		result := []any{}
		for i := 0; i < field.Len(); i++ {
			result = append(result, s.getFieldValue(field.Index(i), ""))
		}
		return result
	} else if field.Kind() == reflect.Int {
		return field.Int()
	} else if field.Kind() == reflect.Map {
		return s.getMapValue(field)
	} else if field.Kind() == reflect.Interface {
		return s.getFieldValue(reflect.ValueOf(field.Interface()), "")
	}

	s.log("== UNHANDLED", field.Kind(), field.Interface())
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
	s.log("Type: ", reflectKind)
	switch reflectKind {
	case reflect.Struct:
		return s.storeStruct(reflectVal.Interface())
	case reflect.Map:
		return s.getMapValue(reflectVal), nil
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
