package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	Id    int `json:"id"`
	Title string
}

var productLIst []Product

func productGet(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(productLIst)

}

func main() {
	mux := http.NewServeMux() // router
	mux.Handle("GET /product", http.HandlerFunc(productGet))
	allowedOrigins := []string{"*"}
	handler := CORSMiddleware(allowedOrigins)(mux)
	fmt.Println("server is running on :3000")    //route
	err := http.ListenAndServe(":3000", handler) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}

}

func init() {
	pro1 := Product{
		Id:    1,
		Title: "Hello",
	}
	pro2 := Product{
		Id:    2,
		Title: "Hello",
	}
	productLIst = append(productLIst, pro1)
	productLIst = append(productLIst, pro2)
}

// CORSMiddleware applies CORS headers to all requests globally
func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Allow specific origins or "*"
			if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if originAllowed(origin, allowedOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Handle OPTIONS preflight
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Helper: check if origin is allowed
func originAllowed(origin string, allowed []string) bool {
	for _, o := range allowed {
		if strings.EqualFold(o, origin) {
			return true
		}
	}
	return false
}
