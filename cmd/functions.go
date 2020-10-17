package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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
