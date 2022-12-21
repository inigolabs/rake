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

		var skip, keep, omit bool
		for _, v := range strings.Split(tField.Tag.Get("writer"), ",") {
			switch v {
			case "readonly", "keep":
				if current == nil {
					if s.input == nil {
						if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
							omit = true
							break
						}
						LoadSources(&current, YamlFileSource(s.filepath))
					} else { // for unit tests
						yaml.Unmarshal(s.input, &current)
					}
				}
				keep = true
			case "omit":
				omit = true
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

		if keep {
			var tag string
			tags := tField.Tag.Get("yaml")
			if len(tags) == 0 {
				tag = strings.ToLower(tField.Name)
			} else if strings.Contains(tags, ",") {
				tag = strings.SplitN(tags, ",", 2)[0]
			} else {
				tag = tags
			}

			rVal := reflect.ValueOf(current[tag])
			if rVal.IsValid() {
				// Set value to current
				vField.Set(rVal)
			} else {
				omit = true
			}
		}

		if omit {
			// Set value to empty
			vField.Set(reflect.Zero(vField.Type()))
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
