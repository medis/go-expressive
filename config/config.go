package config

import (
	module "github.com/medis/go-expressive/internal"
	"github.com/medis/go-expressive/src/App"
)

type Config struct {
	Apps   []*module.App
	Routes []module.Route
}

func Load() *Config {
	var apps []*module.App

	apps = append(
		apps,
		App.NewApp(),
	)

	routes := mergeConfigs(apps)
	return &Config{
		Apps:   apps,
		Routes: routes,
	}
}

func mergeConfigs(apps []*module.App) []module.Route {
	var routes []module.Route
	for _, app := range apps {
		routes = append(routes, app.Module.Routes...)
	}
	return routes
}
