package App

import (
	module "github.com/medis/go-expressive/internal"
	"github.com/medis/go-expressive/src/App/Handlers"
)

func NewApp() *module.App {
	return &module.App{
		Module: module.Module{
			Dependencies: getDependencies(),
			Routes:       getRoutes(),
		},
	}
}

func getDependencies() []module.Dependency {
	return []module.Dependency{
		{
			Invokables: "aa",
			Factories:  "bb",
		},
	}
}

func getRoutes() []module.Route {
	return []module.Route{
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
