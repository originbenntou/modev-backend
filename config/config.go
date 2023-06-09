package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

//go:embed env/default.toml
var defaultEnv []byte

//go:embed env/local.toml
var localEnv []byte

//go:embed env/dev.toml
var devEnv []byte

//go:embed env/prd.toml
var prdEnv []byte

// Config represents configuration root.
type Config struct {
	Env      string
	HTTP     HTTP `toml:"http"`
	AWS      AWS  `toml:"aws"`
	Database Database
}

var config *Config
var once sync.Once

// GetConfig is a function to get Configuration. This loading process occurs once in boot.
func GetConfig() *Config {
	once.Do(func() {
		cfg, err := NewConfig()
		if err != nil {
			panic(err)
		}
		config = cfg
	})

	return config
}

// NewConfig is a function to init and load Configuration from file or environment variables.
func NewConfig() (*Config, error) {
	conf := &Config{}
	conf.Env = "local"

	loadFromEnv(conf)

	err := loadFromToml(defaultEnv, conf)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load env")
	}

	err = loadFromTomlEnv(conf)
	if err != nil {
		return nil, errors.Wrapf(err, "fail loadFromToml env=%v", conf.Env)
	}

	loadDatabaseConfig(conf)

	return conf, nil
}

func loadFromToml(tml []byte, conf *Config) error {
	_, err := toml.NewDecoder(bytes.NewBuffer(tml)).Decode(conf)
	if err != nil {
		return errors.Wrap(err, "fail to decode toml")
	}

	return nil
}

func loadFromTomlEnv(conf *Config) error {
	switch conf.Env {
	case "local":
		return loadFromToml(localEnv, conf)
	case "dev":
		return loadFromToml(devEnv, conf)
	case "prd":
		return loadFromToml(prdEnv, conf)
	default:
		return nil
	}
}

func loadFromEnv(conf *Config) {
	env := os.Getenv("ENV")
	if env != "" {
		conf.Env = env
	}

	fmt.Printf("conf.Env: %v\n", conf.Env)
}
