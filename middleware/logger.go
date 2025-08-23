package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Println("Logger:>>", r.Method, r.URL.Path, time.Since(startTime))
		next.ServeHTTP(w, r)
	})
}
