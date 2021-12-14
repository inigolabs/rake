package rake

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvSource(t *testing.T) {
	type cfg struct {
		StrEmpty     string
		StrPrev      string
		StrVal       string `env:"STR_VAL"`
		IntEmpty     int
		IntPrev      int
		IntVal       int `env:"INT_VAL"`
		BoolEmpty    bool
		BoolPrev     bool
		BoolValTrue  bool `env:"BOOL_VAL_TRUE"`
		BoolValFalse bool `env:"BOOL_VAL_FALSE"`
	}
	actualCfg := cfg{
		StrPrev:  "prev",
		IntPrev:  123,
		BoolPrev: true,
	}

	source := EnvSource("TESTPREFIX")
	os.Setenv("TESTPREFIX_STR_VAL", "v")
	os.Setenv("TESTPREFIX_INT_VAL", "1")
	os.Setenv("TESTPREFIX_BOOL_VAL_TRUE", "true")
	os.Setenv("TESTPREFIX_BOOL_VAL_FALSE", "false")
	source.Load(&actualCfg)

	expectedCfg := cfg{
		StrEmpty:     "",
		StrPrev:      "prev",
		StrVal:       "v",
		IntEmpty:     0,
		IntPrev:      123,
		IntVal:       1,
		BoolEmpty:    false,
		BoolPrev:     true,
		BoolValTrue:  true,
		BoolValFalse: false,
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
