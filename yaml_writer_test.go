package rake

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYamlWriter(t *testing.T) {
	type cfg struct {
		Str      string `writer:"omit"`
		StrD     string
		EmptyStr string
		Int      int64   `writer:"omit"`
		Bool     bool    `writer:"omit"`
		Ptr      *string `writer:"omit"`
		Struct   struct {
			Str string
		}
		StructPtr *struct {
			Str *string
		}
	}

	s := "something"
	actual := cfg{
		Str:  "prev",
		StrD: "asmth",
		Int:  123,
		Bool: true,
		Ptr:  &s,
		Struct: struct{ Str string }{
			Str: "Data",
		},
		StructPtr: &struct{ Str *string }{
			Str: &s,
		},
	}

	expected := actual

	var result bytes.Buffer
	ymw := &yamlWriter{output: &result}
	ymw.Load(&actual)

	require.YAMLEq(t, `str: ""
strd: asmth
int: 0
bool: false
emptystr: ""
ptr: null
struct:
  str: Data
structptr:
  str: something
`, result.String())

	require.EqualValues(t, expected, actual)
}
