package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func TestUpper(t *testing.T) {
	testFunc(t, "upper", upper, "{{ upper . }}", "Foo", "FOO")
}

func TestLower(t *testing.T) {
	testFunc(t, "lower", lower, "{{ lower . }}", "Foo", "foo")
	testFunc(t, "lower", lower, "{{ lower . }}", "foo", "foo")
	testFunc(t, "lower", lower, "{{ lower . }}", 42, "")
}

func testFunc(t *testing.T, name string, f interface{}, tmpl string, input interface{}, expected string) {
	tm := template.New("test")

	tm.Funcs(map[string]interface{}{
		name: f,
	})

	tm, err := tm.Parse(tmpl)
	require.NoError(t, err)

	buffer := bytes.NewBuffer(nil)

	err = tm.Execute(buffer, input)
	require.NoError(t, err)

	assert.Equal(t, expected, buffer.String())
}
