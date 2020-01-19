package module

import "net/http"

type Module struct {
	Dependencies []Dependency
	Routes       []Route
}

type App struct {
	Module Module
}

type Route struct {
	Path string
	Name string
	//Middlewares []interface{}
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

type Dependency struct {
	Invokables string
	Factories  string
}
