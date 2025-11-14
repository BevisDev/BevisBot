package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BevisDev/godev/config"
)

var AppConfig AppConfiguration

func LoadAppConfig() error {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return errors.New("CONFIG_PATH is empty")
	}

	cf := &config.Config{
		Path:       path,
		ConfigType: "yml",
		Dest:       &AppConfig,
	}

	// get profile
	profile := os.Getenv("GO_PROFILE")
	if profile == "" {
		profile = "dev"
	}
	cf.Profile = profile

	// read env
	if profile != "dev" {
		cf.AutoEnv = true
	}

	// load config
	if err := config.NewConfig(cf); err != nil {
		return fmt.Errorf("error load config %v", err)
	}

	var serverCf = AppConfig.Server

	log.Println("================================")
	log.Printf("Load configuration profile %s success", serverCf.Profile)
	log.Printf("Welcome to %s version %s ", serverCf.Name, serverCf.Version)
	log.Println("================================")
	return nil
}
