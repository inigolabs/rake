package rake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlSource(t *testing.T) {
	type cfg struct {
		StrEmpty     string
		StrPrev      string
		StrVal       string `yaml:"StrVal"`
		IntEmpty     int
		IntPrev      int
		IntVal       int `yaml:"IntVal"`
		BoolEmpty    bool
		BoolPrev     bool
		BoolValFalse bool `yaml:"bool_val_false"`
		BoolValTrue  bool `yaml:"bool_val_true"`
	}
	actualCfg := cfg{
		StrPrev:  "prev",
		IntPrev:  123,
		BoolPrev: true,
	}
	source := YamlFileSource("yaml_source_test_data.yml")
	source.Load(&actualCfg)

	expectedCfg := cfg{
		StrEmpty:     "",
		StrPrev:      "prev",
		StrVal:       "val",
		IntEmpty:     0,
		IntPrev:      123,
		IntVal:       5,
		BoolEmpty:    false,
		BoolPrev:     true,
		BoolValFalse: false,
		BoolValTrue:  true,
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
