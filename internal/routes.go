package expressive

import (
	"github.com/gorilla/mux"
)

func (e *Expressive) RegisterRoutes(r *mux.Router) {
	for _, module := range e.Config.Modules {
		for _, route := range module.GetRoutes() {
			r.HandleFunc(route.Path, route.Handler).
				Methods(route.Method)
		}
	}
}
