package Handlers

import (
	"fmt"
	logwrapper "github.com/medis/go-expressive/internal/Logwrapper"
	"net/http"
)

type pingHandlerStruct struct{}

func NewPingHandler() *pingHandlerStruct {
	return &pingHandlerStruct{}
}

func (handler *pingHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logwrapper.GetLogEntry(r).Errorf("hi from ping")
	fmt.Fprintf(w, "pong")
}
