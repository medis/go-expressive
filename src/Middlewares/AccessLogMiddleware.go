package Middlewares

import (
	"log"
	"net/http"
)

type accessLogMiddleware struct{}

func NewAccessLogMiddleware() *accessLogMiddleware {
	return &accessLogMiddleware{}
}

func (middleware *accessLogMiddleware) Handler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
