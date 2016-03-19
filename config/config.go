package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config contains general configuration details
type Config struct {
	Token string `yml:"token"`
	City  string `yml:"city"`
}

// GeneralConfig reads the configuration file and parses its general information
func GeneralConfig() Config {
	configFile := getConfigFile()
	config := Config{}

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

// getConfigFile retrieves the contents of the config file as a byte array
func getConfigFile() []byte {
	filename, err := filepath.Abs("./config.yml")

	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return yamlFile
}
