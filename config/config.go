package config

import (
	AppInterface "github.com/medis/go-expressive/internal/App"
	"github.com/medis/go-expressive/src/App"
)

type Config struct {
	Modules []AppInterface.AppInterface
}

func Load() *Config {
	var modules []AppInterface.AppInterface

	modules = append(
		modules,
		App.NewApp(),
	)

	return &Config{
		Modules: modules,
	}
}
