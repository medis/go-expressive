package Middlewares

import (
	middleware "github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type accessLogMiddleware struct{}

func NewAccessLogMiddleware() *accessLogMiddleware {
	return &accessLogMiddleware{}
}

func (mm *accessLogMiddleware) Handler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestId := middleware.GetReqID(ctx)
			// Do stuff here
			log.Printf("[%s] %s %s\n", requestId, r.RemoteAddr, r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
