package rake

import (
	"reflect"
)

type Source interface {
	Load(configPtr interface{})
}

func LoadSources(cfg interface{}, sources ...Source) {
	if reflect.TypeOf(cfg).Kind() != reflect.Ptr {
		panic("config must be of type struct pointer")
	}

	if len(sources) == 0 {
		panic("no sources passed in")
	}

	for _, source := range sources {
		source.Load(cfg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
