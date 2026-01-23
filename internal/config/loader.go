package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BevisDev/godev/config"
	"github.com/BevisDev/godev/consts"
)

var AppConfig appConfig

func LoadAppConfig() (_ appConfig, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("error loading app config: %v", err)
		}
	}()

	var path = os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config"
	}

	var profile = os.Getenv("GO_PROFILE")
	if profile == "" {
		profile = "dev"
	}

	cf := &config.Config{
		Path:      path,
		Extension: consts.ExtYML,
		Profile:   profile,
	}
	if profile != "dev" {
		cf.AutoEnv = true
	}
	resp := config.MustLoad[appConfig](cf)
	if resp.Data == nil {
		return AppConfig, errors.New("data empty")
	}
	
	AppConfig = *resp.Data
	return AppConfig, nil
}
