package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

const (
	propertiesTag = "properties"
	omitemptyTag  = "omitempty"
)

func Serialize(person Person) string {
	strBuild := strings.Builder{}
	typ := reflect.TypeOf(person)
	val := reflect.ValueOf(person)

	for i := 0; i < typ.NumField(); i++ {
		property, omit := parsePropertiesTag(typ.Field(i))
		if omit && checkFieldZeroValue(val.Field(i)) {
			continue
		}

		strBuild.WriteString(property)
		strBuild.WriteString("=")
		strBuild.WriteString(parseValueToString(val.Field(i)))
		if i != typ.NumField()-1 {
			strBuild.WriteString("\n")
		}
	}

	return strBuild.String()
}

func parsePropertiesTag(field reflect.StructField) (string, bool) {
	var omitempty bool
	tag := field.Tag.Get(propertiesTag)
	if tag == "" {
		return "", false
	}

	parts := strings.Split(tag, ",")
	if len(parts) > 1 && parts[1] == omitemptyTag {
		omitempty = true
	}

	return parts[0], omitempty
}

func parseValueToString(val reflect.Value) string {
	switch val.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.Itoa(int(val.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.Itoa(int(val.Uint()))
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, val.Type().Bits())
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case reflect.Struct:
		return parseValueToString(val)
	default:
		panic(fmt.Sprintf("unhandled default case for type: %s", val.Type().Kind()))
	}
}

func checkFieldZeroValue(val reflect.Value) bool {
	return val.IsZero()
}
