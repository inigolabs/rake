package rake

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

func YamlFileWriter(filepath string) *yamlWriter {
	return &yamlWriter{
		filepath: filepath,
	}
}

type yamlWriter struct {
	filepath string
	output   *bytes.Buffer
}

func (s *yamlWriter) Load(configPtr interface{}) {
	t := reflect.TypeOf(configPtr)
	v := reflect.ValueOf(configPtr)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)

		fieldTags := tField.Tag.Get("writer")
		if !strings.Contains(fieldTags, "omit") {
			continue
		}

		if vField.IsZero() || !vField.IsValid() {
			continue
		}

		// Copy addressable value
		dst := reflect.New(vField.Type())
		dst.Elem().Set(vField)

		// Reset to prior value
		defer vField.Set(dst.Elem())

		// Set empty to empty
		vField.Set(reflect.Zero(vField.Type()))
	}

	var err error
	var result io.Writer = s.output

	if s.output == nil {
		result, err = os.Create(s.filepath)
		check(err)
	}

	encoder := yaml.NewEncoder(result)
	err = encoder.Encode(configPtr)
	check(err)
}
