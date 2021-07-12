package rake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalMap(t *testing.T) {
	type CfgSub struct {
		Str  string
		Int  int
		Bool bool
	}
	type Cfg struct {
		Sub CfgSub
	}

	var actualCfg Cfg

	source := LocalMap(map[string]interface{}{
		"Sub.Str":  "string1",
		"Sub.Int":  1,
		"Sub.Bool": true,
	})
	source.Load(&actualCfg)

	expectedCfg := Cfg{
		Sub: CfgSub{
			Str:  "string1",
			Int:  1,
			Bool: true,
		},
	}

	assert.Equal(t, expectedCfg, actualCfg)
}
