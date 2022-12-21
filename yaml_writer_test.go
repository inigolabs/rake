package rake

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestYamlWriterOmit(t *testing.T) {
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

func TestYamlWriterKeep(t *testing.T) {
	type cfg struct {
		Str        string `writer:"keep"`
		StrCustom  string `yaml:"abc" writer:"keep"`
		StrCustom2 string `yaml:"cba,omitempty" writer:"keep"`
		Int        int    `writer:"readonly"`
		Empty      string `yaml:"Version,omitempty" writer:"keep"`
	}

	input := cfg{
		Str:        "Something",
		StrCustom:  "Something2",
		StrCustom2: "Something3",
		Int:        1,
	}
	out, err := yaml.Marshal(input)
	require.NoError(t, err)

	modified := input
	modified.Str = "Stuff"
	modified.StrCustom = "stuff"
	modified.StrCustom2 = "stuff"
	modified.Int = 12342
	modified.Empty = "NOT EMPTY"

	var output bytes.Buffer
	ymw := &yamlWriter{output: &output, input: out}
	ymw.Load(&modified)

	var result cfg
	yaml.Unmarshal(output.Bytes(), &result)

	require.EqualValues(t, input, result)
}
