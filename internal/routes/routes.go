package routes

import (
	"github.com/gorilla/mux"
	"github.com/medis/go-expressive/config"
)

func RegisterRoutes(r *mux.Router, config *config.Config) {
	for _, route := range config.Routes {
		r.HandleFunc(route.Path, route.Handler).
			Methods(route.Method)
	}
}
