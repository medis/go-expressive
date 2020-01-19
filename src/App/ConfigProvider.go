package App

import (
	"github.com/medis/go-expressive/internal/Route"
	"github.com/medis/go-expressive/src/App/Handlers"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

//func getDependencies() []config.Dependency {
//	return []config.Dependency{
//		{
//			Invokables: "aa",
//			Factories:  "bb",
//		},
//	}
//}

func (a *App) GetRoutes() []Route.Route {
	return []Route.Route{
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
