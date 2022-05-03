package rake

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSource(t *testing.T) {
	type cfg struct {
		StrEmpty     string
		StrPrev      string
		StrVal       string `default:"v"`
		StrColonVal  string `default:"a:b"`
		IntEmpty     int
		IntPrev      int
		IntVal       int `default:"1"`
		BoolEmpty    bool
		BoolPrev     bool
		BoolValFalse bool          `default:"false"`
		BoolValTrue  bool          `default:"true"`
		Duration     time.Duration `default:"3s"`
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
		StrColonVal:  "a:b",
		IntEmpty:     0,
		IntPrev:      123,
		IntVal:       1,
		BoolEmpty:    false,
		BoolPrev:     true,
		BoolValFalse: false,
		BoolValTrue:  true,
		Duration:     time.Second * 3,
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
