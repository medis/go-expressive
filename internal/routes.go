package expressive

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).
			Methods(route.Method)
	}
}
