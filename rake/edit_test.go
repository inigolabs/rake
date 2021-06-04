package rake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterStruct(t *testing.T) {
	type Config struct {
		A int
		B string
		C bool
		D struct {
			DA int
			DB string
			DC bool
		}
	}
	cfg := Config{}

	var actual []Path
	editFunc := func(attrPtr interface{}, path Path, tags map[string]string) {
		actual = append(actual, path)
	}

	IterConfig(&cfg, editFunc)

	expected := []Path{
		[]string{"A"},
		[]string{"B"},
		[]string{"C"},
		[]string{"D", "DA"},
		[]string{"D", "DB"},
		[]string{"D", "DC"},
	}

	assert.Equal(t, expected, actual)
}
