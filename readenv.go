package readenv

import (
	"os"
	"reflect"
)

func Read(s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Pointer {
		panic("s is not a pointer")
	}

	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		panic("*s is not a struct")
	}

	for i := 0; i < elem.NumField(); i += 1 {
		fieldValue := elem.Field(i)
		fieldType := elem.Type().Field(i)

		if fieldValue.Kind() != reflect.String {
			panic("Every field type must be string")
		}

		envname := fieldType.Tag.Get("env")
		if envname == "" {
			panic("Environment variable name should not be empty")
		}

		env, ok := os.LookupEnv(envname)
		if !ok {
			panic("Environment variable `" + envname + "` is not found")
		}

		fieldValue.Set(reflect.ValueOf(env))
	}
}
