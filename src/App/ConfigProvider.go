package App

import (
	expressive "github.com/medis/go-expressive/internal"
	"github.com/medis/go-expressive/src/App/Handlers"
)

func NewApp() *expressive.App {
	return &expressive.App{
		Module: expressive.Module{
			Dependencies: getDependencies(),
			Routes:       getRoutes(),
		},
	}
}

func getDependencies() []expressive.Dependency {
	return []expressive.Dependency{
		{
			Invokables: "aa",
			Factories:  "bb",
		},
	}
}

func getRoutes() []expressive.Route {
	return []expressive.Route{
		{
			Path:    "/",
			Name:    "home",
			Handler: Handlers.Home,
			Method:  "GET",
		},
		{
			Path:    "/ping",
			Name:    "ping",
			Handler: Handlers.Ping,
			Method:  "GET",
		},
	}
}
