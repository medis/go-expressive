package Handlers

import (
	"fmt"
	"net/http"
)

type pingHandlerStruct struct{}

func NewPingHandler() *pingHandlerStruct {
	return &pingHandlerStruct{}
}

func (handler *pingHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
