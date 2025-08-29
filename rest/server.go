package rest

import (
	"ecommace/config"
	"ecommace/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	warpedMux := manager.WrapMux(mux)
	InitRoute(mux, manager)
	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server is running on", addr)   //route
	err := http.ListenAndServe(addr, warpedMux) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}
}
