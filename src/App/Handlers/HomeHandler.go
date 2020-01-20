package Handlers

import (
	"fmt"
	"net/http"
)

type homeHandlerStruct struct{}

func NewHomeHandler() *homeHandlerStruct {
	return &homeHandlerStruct{}
}

func (handler *homeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
