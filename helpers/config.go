package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ReadInConfig function to handle reading config file with format YAML
func ReadInConfig[T any](path string) (T, error) {
	var t T
	bytes, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &t); nil != err {
		return nil, err
	}

	return t, nil
}
