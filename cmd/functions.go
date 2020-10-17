package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
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

func toSnake(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strcase.ToSnake(input.String())
	}

	return ""
}

func toKebab(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strcase.ToKebab(input.String())
	}

	return ""
}

func toCamel(input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strcase.ToCamel(input.String())
	}

	return ""
}

func toDelimited(delimiter reflect.Value, input reflect.Value) string {
	if input.Kind() == reflect.String {
		return strcase.ToDelimited(input.String(), delimiter.String()[0])
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

func toSha(input reflect.Value) string {
	if input.Kind() == reflect.String {
		h := sha1.New()
		h.Write([]byte(input.String()))

		s := h.Sum(nil)

		return hex.EncodeToString(s)
	}

	return ""
}
func toSha256(input reflect.Value) string {
	if input.Kind() == reflect.String {
		h := sha256.New()
		h.Write([]byte(input.String()))

		s := h.Sum(nil)

		return hex.EncodeToString(s)
	}

	return ""
}

func toSha512(input reflect.Value) string {
	if input.Kind() == reflect.String {
		h := sha512.New()
		h.Write([]byte(input.String()))

		s := h.Sum(nil)

		return hex.EncodeToString(s)
	}

	return ""
}
