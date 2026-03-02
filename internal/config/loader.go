package config

import (
	"errors"
	"os"

	"github.com/BevisDev/godev/config"
	"github.com/BevisDev/godev/consts"
)

type Config struct {
	appConfig *appConfig
}

func New() (*Config, error) {
	appConfig, err := load()
	if err != nil {
		return nil, err
	}
	return &Config{
		appConfig: appConfig,
	}, nil
}

func load() (*appConfig, error) {
	profile := os.Getenv("GO_PROFILE")
	if profile == "" {
		profile = "dev"
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config"
	}

	autoEnv := false
	if profile != "dev" {
		autoEnv = true
	}

	cf := &config.Config{
		Path:    path,
		Ext:     consts.ExtYML,
		AutoEnv: autoEnv,
		Profile: profile,
	}

	resp, err := config.Load[*appConfig](cf)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, errors.New("no data")
	}

	return resp.Data, nil
}

func (c *Config) AppConfig() appConfig {
	return *c.appConfig
}
