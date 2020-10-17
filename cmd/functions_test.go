package cmd

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringFunctions(t *testing.T) {
	testTempl(t, "upper", toUpper, "{{ upper . }}", "Foo", "FOO")
	testTempl(t, "lower", toLower, "{{ lower . }}", "Foo", "foo")

	testTempl(t, "suffix", toSuffix, `{{ suffix "-san" . }}`, "Foo", "Foo-san")
	testTempl(t, "prefix", toPrefix, `{{ prefix "Mr " . }}`, "Foo", "Mr Foo")

	testTempl(t, "kebab", toKebab, "{{ kebab . }}", "Foo bar", "foo-bar")
	testTempl(t, "camel", toCamel, "{{ camel . }}", "Foo bar", "FooBar")
	testTempl(t, "snake", toSnake, "{{ snake . }}", "Foo bar", "foo_bar")
	testTempl(t, "delimited", toDelimited, `{{ delimited "." . }}`, "Foo bar", "foo.bar")
	testTempl(t, "delimited", toDelimited, `{{ . | delimited "." }}`, "Foo bar", "foo.bar")
}

func TestHashFunctions(t *testing.T) {
	testTempl(t, "md5", toMd5, "{{ md5 . }}", "Foo", "1356c67d7ad1638d816bfb822dd2c25d")
	testTempl(t, "sha", toSha, "{{ sha . }}", "Foo", "201a6b3053cc1422d2c3670b62616221d2290929")
	testTempl(t, "sha256", toSha256, "{{ sha256 . }}", "Foo", "1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa")
	testTempl(t, "sha512", toSha512, "{{ sha512 . }}", "Foo", "4abcd2639957cb23e33f63d70659b602a5923fafcfd2768ef79b0badea637e5c837161aa101a557a1d4deacbd912189e2bb11bf3c0c0c70ef7797217da7e8207")
}

func testTempl(t *testing.T, name string, f interface{}, tmpl string, input interface{}, expected string) {
	t.Run(name, func(t *testing.T) {
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
	})
}
