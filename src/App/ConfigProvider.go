package App

import (
	"github.com/medis/go-expressive/di"
	"github.com/medis/go-expressive/internal/Route"
	"github.com/medis/go-expressive/internal/Template"
	"github.com/medis/go-expressive/src/App/Handlers"
	"github.com/medis/go-expressive/src/Middlewares"
)

type App struct {
	template *Template.Template
}

func NewApp(template *Template.Template) *App {
	return &App{
		template: template,
	}
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
			Handler: Handlers.NewHomeHandler(a.template),
			Middlewares: Route.Middlewares{
				Middlewares.NewAccessLogMiddleware().Handler(),
			},
			Method: "GET",
		},
		{
			Path:    "/ping",
			Name:    "ping",
			Handler: di.InitialisePingHandler(),
			Method:  "GET",
		},
	}
}
