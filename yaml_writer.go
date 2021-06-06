package rake

import (
	"os"

	"gopkg.in/yaml.v2"
)

func YamlFileWriter(filepath string) *yamlWriter {
	return &yamlWriter{
		filepath: filepath,
	}
}

type yamlWriter struct {
	filepath string
}

func (s *yamlWriter) Load(configPtr interface{}) {
	file, err := os.Create(s.filepath)
	check(err)
	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(configPtr)
	check(err)
}
