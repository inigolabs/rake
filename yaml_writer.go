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
	input    []byte
}

func (s *yamlWriter) Load(configPtr interface{}) {
	var current map[string]interface{}

	t := reflect.TypeOf(configPtr)
	v := reflect.ValueOf(configPtr)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)

		var skip bool
		for _, v := range strings.Split(tField.Tag.Get("writer"), ",") {
			switch v {
			case "omit":
				if current == nil {
					if s.input == nil {
						if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
							break
						}
						LoadSources(&current, YamlFileSource(s.filepath))
					} else { // for unit tests
						yaml.Unmarshal(s.input, &current)
					}
				}
			default:
				skip = true
			}
		}

		if skip || vField.IsZero() || !vField.IsValid() {
			continue
		}

		// Copy addressable value
		dst := reflect.New(vField.Type())
		dst.Elem().Set(vField)

		// Reset to prior value
		defer vField.Set(dst.Elem())

		tag := tField.Tag.Get("yaml") // yaml:"something"
		if len(tag) == 0 {
			tag = strings.ToLower(tField.Name) // yaml:""
		} else if strings.Contains(tag, ",") {
			tag = strings.SplitN(tag, ",", 2)[0] // yaml:"something,omitempty"
		}

		rVal := reflect.ValueOf(current[tag])
		if rVal.IsValid() {
			vField.Set(rVal) // Set value to current
		} else {
			vField.Set(reflect.Zero(vField.Type())) // set value to empty when it's invalid
		}
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
