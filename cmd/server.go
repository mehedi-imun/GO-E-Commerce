package cmd

import (
	"ecommace/database"
	"ecommace/middleware"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux() // router
	mux.Handle("GET /products", http.HandlerFunc(database.ProductGet))
	mux.Handle("GET /products/{id}", http.HandlerFunc(database.GetProductByID))
	allowedOrigins := []string{"*"}
	handler := middleware.CORSMiddleware(allowedOrigins)(mux)
	fmt.Println("server is running on :3000")    //route
	err := http.ListenAndServe(":3000", handler) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}
}
