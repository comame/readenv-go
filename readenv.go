package readenv

import (
	"os"
	"reflect"
	"strings"
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

		tag := parseTag(fieldType.Tag.Get("env"))

		env, ok := os.LookupEnv(tag.envname)
		if !ok && tag.opt != option(optional) {
			panic("Environment variable `" + tag.envname + "` is not found")
		}

		fieldValue.Set(reflect.ValueOf(env))
	}
}

type option int

const (
	none int = iota
	optional
)

type tag struct {
	envname string
	opt     option
}

func parseTag(t string) tag {
	s := strings.Split(t, ",")

	if !(len(s) == 1 || len(s) == 2) {
		panic("env tag `" + t + "` is invalid form")
	}

	envname := s[0]
	opt := option(none)

	if len(s) == 2 {
		switch s[1] {
		case "optional":
			opt = option(optional)
		}
	}

	envname = strings.TrimSpace(envname)
	if envname == "" {
		panic("envname must not be empty string")
	}

	return tag{
		envname: envname,
		opt:     opt,
	}
}
