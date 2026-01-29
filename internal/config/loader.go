package config

import (
	"errors"
	"github.com/BevisDev/godev/config"
	"github.com/BevisDev/godev/consts"
	"os"
)

var AppConfig appConfig

func Load() error {
	profile := os.Getenv("PROFILE")
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

	resp, err := config.Load[appConfig](cf)
	if err != nil {
		return err
	}

	if resp.Data == nil {
		return errors.New("no data")
	}

	AppConfig = *resp.Data
	return nil
}
