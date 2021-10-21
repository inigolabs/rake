package rake

import (
	"reflect"
	"strings"
)

type Path []string

type EditStructFieldFunc func(attrPtr interface{}, path Path, tags map[string]string)

func IterConfig(cfg interface{}, editFunc EditStructFieldFunc) {
	iterStruct(cfg, Path{}, editFunc)
}

func iterStruct(elem interface{}, path Path, editFunc EditStructFieldFunc) {
	t := reflect.TypeOf(elem)
	v := reflect.ValueOf(elem)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)
		currentPath := append(path, tField.Name)

		if tField.Type.Kind() == reflect.Struct {
			iterStruct(vField.Addr().Interface(), currentPath, editFunc)
			continue
		}

		if tField.Type.Kind() == reflect.Ptr && tField.Type.Elem().Kind() == reflect.Struct {
			vField.Set(reflect.New(tField.Type.Elem()))
			iterStruct(vField.Interface(), currentPath, editFunc)
			continue
		}

		structTags := make(map[string]string)
		if tField.Tag != "" {
			tagPairs := strings.Split(string(tField.Tag), " ")
			for _, pair := range tagPairs {
				tagKeyVal := strings.SplitN(pair, ":", 2)
				key := tagKeyVal[0]
				val := strings.Trim(tagKeyVal[1], "\"")
				structTags[key] = val
			}
		}
		editFunc(vField.Addr().Interface(), currentPath, structTags)
	}
}
