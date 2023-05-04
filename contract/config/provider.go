package config

import (
	"errors"
	"os"

	"github.com/zackyymughnii/runtime-go/contract"
	"gopkg.in/yaml.v3"
)

type Container[T any] interface {
	GetConfigPath() string
	SetConfig(cfg T)
}

func Provider[T any]() contract.Provide {
	return func(arg any) error {
		container, ok := arg.(Container[T])
		if !ok {
			return errors.New("application doesn't implement config container")
		}

		path := container.GetConfigPath()
		p, err := os.ReadFile(path)
		if nil != err {
			return err
		}

		var cfg T
		if err = yaml.Unmarshal(p, &cfg); nil != err {
			return err
		}

		container.SetConfig(cfg)

		return nil
	}
}
