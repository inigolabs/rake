package rake

import "github.com/caarlos0/env/v9"

func EnvSource(prefix string) *envSource {
	return &envSource{
		prefix: prefix,
	}
}

type envSource struct {
	prefix string
}

func (s *envSource) Load(configPtr interface{}) {
	if err := env.ParseWithOptions(configPtr, env.Options{
		Prefix: s.prefix + "_",
	}); err != nil {
		panic(err) //
	}
}
