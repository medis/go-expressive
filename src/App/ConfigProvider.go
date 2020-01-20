package App

import (
	"github.com/medis/go-expressive/internal/Route"
	"github.com/medis/go-expressive/src/App/Handlers"
	"github.com/medis/go-expressive/src/Middlewares"
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
			Handler: Handlers.NewHomeHandler(),
			Middlewares: Route.Middlewares{
				Middlewares.NewAccessLogMiddleware().Handler(),
			},
			Method: "GET",
		},
		{
			Path:    "/ping",
			Name:    "ping",
			Handler: Handlers.NewPingHandler(),
			Method:  "GET",
		},
	}
}
