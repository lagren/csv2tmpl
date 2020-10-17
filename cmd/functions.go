package cmd

import (
	"reflect"
	"strings"
)

func upper(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strings.ToUpper(input.String())
	}

	return ""
}

func lower(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strings.ToLower(input.String())
	}

	return ""
}
