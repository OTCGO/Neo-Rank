package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

type (
	// Config top level struct representing the config
	// for the node.
	Config struct {
		ApplicationConfiguration
	}

	// ProtocolConfiguration represents the protolcol config.
	ProtocolConfiguration struct {
	}

	// ApplicationConfiguration config specific to the node.
	ApplicationConfiguration struct {
		Environment string `yaml:"Environment"`
	}
)

var config Config

// Loadattempts to load the config from the give
// path and env.
func Load(path string, env string) (Config, error) {
	configPath := fmt.Sprintf("%s/%s.yml", path, env)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, errors.Wrap(err, "Unable to load config")
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, errors.Wrap(err, "Unable to read config")
	}

	config = Config{
		ApplicationConfiguration: ApplicationConfiguration{},
	}

	err = yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		return Config{}, errors.Wrap(err, "Problem unmarshaling config json data")
	}

	return config, nil
}
