package cmd

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpper(t *testing.T) {
	testTempl(t, "upper", toUpper, "{{ upper . }}", "Foo", "FOO")
}

func TestLower(t *testing.T) {
	testTempl(t, "lower", toLower, "{{ lower . }}", "Foo", "foo")
	testTempl(t, "lower", toLower, "{{ lower . }}", "foo", "foo")
	testTempl(t, "lower", toLower, "{{ lower . }}", 42, "")
}

func TestMD5(t *testing.T) {
	testTempl(t, "md5", toMd5, "{{ md5 . }}", "Foo", "1356c67d7ad1638d816bfb822dd2c25d")
}

func testTempl(t *testing.T, name string, f interface{}, tmpl string, input interface{}, expected string) {
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
