package cmd

import (
	"ecommace/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)

	mux := http.NewServeMux()
	InitRoute(mux, manager)
	allowedOrigins := []string{"*"}
	handler := middleware.CORSMiddleware(allowedOrigins)(mux)
	fmt.Println("server is running on :3000")    //route
	err := http.ListenAndServe(":3000", handler) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}
}
