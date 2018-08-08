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
		ApplicationConfiguration ApplicationConfiguration `yaml:"ApplicationConfiguration"`
		Mongodb                  Mongodb                  `yaml:"Mongodb"`
		Redis                    Redis                    `yaml:"Redis"`
	}

	// ApplicationConfiguration config specific to the node.
	ApplicationConfiguration struct {
		Environment string   `yaml:"Environment"`
		Nodes       []string `yaml:"Nodes"`
	}

	Mongodb struct {
		Addrs     []string `yaml:"Addrs"`
		Direct    bool     `yaml:"Direct"`
		Database  string   `yaml:"Database"`
		Source    string   `yaml:"Source"`
		Username  string   `yaml:"Username"`
		Password  string   `yaml:"Password"`
		PoolLimit int      `yaml:"PoolLimit"`
	}

	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	}
)

var config Config

// Loadattempts to load the config from the give
// path and env.
func Load(path string, env string) (Config, error) {
	configPath := fmt.Sprintf("%s/%s.yml", path, env)
	fmt.Println("configPath", configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, errors.Wrap(err, "Unable to load config")
	}

	configData, err := ioutil.ReadFile(configPath)

	// fmt.Println("configData", string(configData))

	if err != nil {
		return Config{}, errors.Wrap(err, "Unable to read config")
	}

	config = Config{
		ApplicationConfiguration: ApplicationConfiguration{},
		Mongodb:                  Mongodb{},
		Redis:                    Redis{},
	}

	err = yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		return Config{}, errors.Wrap(err, "Problem unmarshaling config json data")
	}

	fmt.Println("config", config)

	return config, nil
}

func GetConfig() (config Config) {
	if &config != nil {
		errors.New("GetConfig error")
	}
	return config
}
