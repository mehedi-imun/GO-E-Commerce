package cmd

import (
	"ecommace/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	warpedMux := manager.WrapMux(mux)
	InitRoute(mux, manager)
	fmt.Println("server is running on :3000")      //route
	err := http.ListenAndServe(":3000", warpedMux) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}
}
