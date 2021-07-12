package rake

import (
	"fmt"
	"strings"
)

func LocalMap(local map[string]interface{}) *localMap {
	return &localMap{
		localData: local,
	}
}

type localMap struct {
	localData map[string]interface{}
}

func (s *localMap) Load(configPtr interface{}) {

	editFunc := func(attrPtr interface{}, path Path, tags map[string]string) {
		pathStr := strings.Join(path, ".")
		if data, found := s.localData[pathStr]; found {
			switch a := attrPtr.(type) {
			case *string:
				*a = data.(string)
			case *int:
				*a = data.(int)
			case *bool:
				*a = data.(bool)
			default:
				panic(fmt.Errorf("unsupported config attr type %T with path %v", attrPtr, path))
			}
		}
	}

	IterConfig(configPtr, editFunc)
}
