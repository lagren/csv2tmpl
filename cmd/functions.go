package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"strings"
)

func toUpper(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strings.ToUpper(input.String())
	}

	return ""
}

func toLower(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strings.ToLower(input.String())
	}

	return ""
}

func toMd5(input reflect.Value) string {
	if input.Kind() == reflect.String {
		h := md5.New()
		h.Write([]byte(input.String()))

		s := h.Sum(nil)

		return hex.EncodeToString(s)
	}

	return ""
}
