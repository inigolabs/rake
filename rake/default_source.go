package rake

import (
	"fmt"
	"strconv"
)

func DefaultSource() *defaultSource {
	return &defaultSource{}
}

type defaultSource struct{}

func (s *defaultSource) Load(configPtr interface{}) {
	editFunc := func(attrPtr interface{}, path Path, tags map[string]string) {
		if val, found := tags["default"]; found {
			switch a := attrPtr.(type) {
			case *string:
				*a = val
			case *int:
				intVal, err := strconv.Atoi(val)
				check(err)
				*a = intVal
			case *bool:
				boolVal, err := strconv.ParseBool(val)
				check(err)
				*a = boolVal
			default:
				panic(fmt.Errorf("unsupported config attr type %T with path %v", attrPtr, path))
			}
		}
	}

	IterConfig(configPtr, editFunc)
}
