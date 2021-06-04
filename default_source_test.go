package rake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSource(t *testing.T) {
	type cfg struct {
		StrEmpty     string
		StrPrev      string
		StrVal       string `default:"v"`
		IntEmpty     int
		IntPrev      int
		IntVal       int `default:"1"`
		BoolEmpty    bool
		BoolPrev     bool
		BoolValFalse bool `default:"false"`
		BoolValTrue  bool `default:"true"`
	}
	actualCfg := cfg{
		StrPrev:  "prev",
		IntPrev:  123,
		BoolPrev: true,
	}
	source := DefaultSource()
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
		BoolValFalse: false,
		BoolValTrue:  true,
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
