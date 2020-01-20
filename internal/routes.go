package expressive

import (
	"github.com/go-chi/chi"
)

func (e *Expressive) RegisterRoutes(r *chi.Mux) {
	for _, module := range e.Config.Modules {
		for _, route := range module.GetRoutes() {
			r.With(route.Middlewares...).Method(route.Method, route.Path, route.Handler)
		}
	}
}
