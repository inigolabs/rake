package rake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigMultiSource(t *testing.T) {
	type cfg struct {
		StrEmpty   string
		StrPrev    string
		StrDefault string `default:"def"`
		StrVal     string `default:"def" yaml:"StrVal"`
		IntEmpty   int
		IntPrev    int
		IntDefault int `default:"1"`
		IntVal     int `default:"2" yaml:"IntVal"`
	}
	actualCfg := cfg{
		StrPrev: "prev",
		IntPrev: 123,
	}

	LoadSources(&actualCfg,
		DefaultSource(),
		YamlFileSource("yaml_source_test_data.yml"),
	)

	expectedCfg := cfg{
		StrEmpty:   "",
		StrPrev:    "prev",
		StrDefault: "def",
		StrVal:     "val",
		IntEmpty:   0,
		IntPrev:    123,
		IntDefault: 1,
		IntVal:     5,
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
