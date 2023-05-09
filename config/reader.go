package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// Read config on yaml file and check if values contains variable fields get it on os.Env
func Read[T any](path, key string) (*T, error) {
	keyVal, err := read(path, key)
	if nil != err {
		return nil, err
	}

	keyVal = val(keyVal)

	var cfg T
	if err = mapstructure.Decode(keyVal, &cfg); nil != err {
		panic(err)
	}

	return &cfg, nil
}

func read(path, keys string) (map[string]any, error) {
	f, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	var result = make(map[string]any)
	if err = yaml.NewDecoder(f).Decode(&result); nil != err {
		return nil, err
	}

	return readKeys(keys, result)
}

func readKeys(keys string, arg map[string]any) (map[string]any, error) {
	if keys == "" {
		return arg, nil
	}

	keyArray := strings.Split(keys, "/")
	for _, k := range keyArray {
		v, exists := arg[k]
		if !exists {
			return nil, fmt.Errorf("config with keys %s not found", k)
		}

		mapVal, ok := v.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid format config for key %s", k)
		}

		arg = mapVal
	}

	return arg, nil
}

func val(arg map[string]any) map[string]any {
	for k, v := range arg {
		rawV, ok := v.(map[string]any)
		if ok {
			arg[k] = val(rawV)
			continue
		}

		valString := fmt.Sprintf("%v", v)
		match, _ := regexp.MatchString(`\${(\w+)}`, valString)
		if match {
			keyEnv := regexp.
				MustCompile(`^\${|}$`).
				ReplaceAllString(valString, ``)

			arg[k] = os.Getenv(keyEnv)
		}
	}

	return arg
}
