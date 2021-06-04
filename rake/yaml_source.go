package rake

import (
	"os"

	"gopkg.in/yaml.v2"
)

func YamlFileSource(filepath string) *yamlSource {
	return &yamlSource{
		filepath: filepath,
	}
}

type yamlSource struct {
	filepath string
}

func (s *yamlSource) Load(configPtr interface{}) {
	file, err := os.OpenFile(s.filepath, os.O_RDONLY, os.ModePerm)
	check(err)
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(configPtr)
	check(err)
}
