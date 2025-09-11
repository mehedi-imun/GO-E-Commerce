package rest

import (
	"ecommace/config"
	"ecommace/rest/handlers/product"
	"ecommace/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	cnf            config.Config
}

func NewServer(cnf config.Config,
	productHandler *product.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	warpedMux := manager.WrapMux(mux)

	server.productHandler.Product_Route(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("server is running on", addr)   //route
	err := http.ListenAndServe(addr, warpedMux) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}
}
