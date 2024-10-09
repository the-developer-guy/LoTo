package internals

import (
	"os"

	"gopkg.in/yaml.v3"
)

type LotoConfig struct {
	Name     string      `yaml:"name"`
	Services *[]Services `yaml:"services"`
}

type Services struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func GetConfig() (*LotoConfig, error) {
	configFile, err := os.ReadFile("config.yaml")
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
