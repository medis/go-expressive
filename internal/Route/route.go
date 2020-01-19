package Route

import "net/http"

type Route struct {
	Path string
	Name string
	//Middlewares []interface{}
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}
