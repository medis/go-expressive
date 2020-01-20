package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	AppInterface "github.com/medis/go-expressive/internal/App"
	"github.com/medis/go-expressive/src/App"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Modules []AppInterface.AppInterface
	Server  struct {
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	App struct {
		Secret       string        `yaml:"secret" envconfig:"APP_SECRET"`
		ShutdownTime time.Duration `yaml:"shutdown_timeout" envconfig:"APP_SHUTDOWN_TIMEOUT"`
	}
}

func Load() *Config {
	cfg := Config{}
	// Read config file.
	readConfig(&cfg)
	// Overwrite with env variables.
	readEnv(&cfg)
	// Register modules.
	loadModules(&cfg)

	return &cfg
}

// Load configuration from config file.
func readConfig(cfg *Config) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(fmt.Sprintf("%s/config/config.yml", dir))
	if err != nil {
		panic("Cannot load configuration file")
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic("Cannot decode configuration file")
	}
}

// Load and overwrite config from ENV variables.
func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic("Could not read environment configuration")
	}
}

func loadModules(cfg *Config) {
	cfg.Modules = append(
		[]AppInterface.AppInterface{},
		App.NewApp(),
	)
}
