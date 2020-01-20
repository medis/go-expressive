package Route

import "net/http"

type Middlewares []func(next http.Handler) http.Handler

type Route struct {
	Path        string
	Name        string
	Handler     http.Handler
	Middlewares Middlewares
	Method      string
}
