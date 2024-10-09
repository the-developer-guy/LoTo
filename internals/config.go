package internals

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type LotoConfig struct {
	Name     string     `yaml:"name"`
	Services *[]Service `yaml:"services"`
}

type Service struct {
	Name   string `yaml:"name"`
	Url    string `yaml:"url"`
	Locked bool
}

func GetConfig() (*LotoConfig, error) {
	configLocation := os.Getenv("LOTO_CONFIG_PATH")
	if configLocation == "" {
		return nil, fmt.Errorf("no LOTO_CONFIG_PATH set")
	}
	configFile, err := os.ReadFile(configLocation)
	if err != nil {
		return nil, err
	}
	var config LotoConfig
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
