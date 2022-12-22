package rake

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestYamlWriterOmit(t *testing.T) {
	type cfg struct {
		Int        int    `writer:"omit"`
		Str        string `writer:"omit"`
		StrCustom  string `yaml:"abc" writer:"omit"`
		StrCustom2 string `yaml:"cba,omitempty" writer:"omit"`
		Empty      string `yaml:"Version,omitempty" writer:"omit"`
	}

	input := cfg{
		Int:        1,
		Str:        "Something",
		StrCustom:  "Something2",
		StrCustom2: "Something3",
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
