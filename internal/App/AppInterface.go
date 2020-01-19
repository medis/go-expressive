package AppInterface

import "github.com/medis/go-expressive/internal/Route"

type AppInterface interface {
	GetRoutes() []Route.Route
}
