package middlewares

import (
	"auth-service/logger"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("Got request on '%v' PATH, '%v' METHOD", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}
