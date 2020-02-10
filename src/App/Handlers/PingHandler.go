package Handlers

import (
	"github.com/medis/go-expressive/internal/Response"
	"net/http"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (handler *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//logwrapper.GetLogEntry(r).Errorf("hi from ping")
	headers := make(map[string]string)
	Response.JsonResponse("Pong", w, 200, headers)
}
