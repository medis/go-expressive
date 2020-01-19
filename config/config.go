package config

import (
	expressive "github.com/medis/go-expressive/internal"
	"github.com/medis/go-expressive/src/App"
)

type Config struct {
	Apps   []*expressive.App
	Routes []expressive.Route
}

func Load() *Config {
	var apps []*expressive.App

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

func mergeConfigs(apps []*expressive.App) []expressive.Route {
	var routes []expressive.Route
	for _, app := range apps {
		routes = append(routes, app.Module.Routes...)
	}
	return routes
}
