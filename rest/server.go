package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommace/config"
	"ecommace/rest/handlers/product"
	"ecommace/rest/handlers/user"
	"ecommace/rest/middleware"
)

// Server holds dependencies
type Server struct {
	cnf            *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
}

// NewServer creates a server with injected handlers and config
func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	productHandler *product.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

// Start runs the HTTP server
func (s *Server) Start() {
	// 1️⃣ Create middleware manager
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	// 2️⃣ Create HTTP mux
	mux := http.NewServeMux()

	// 3️⃣ Register routes
	if s.userHandler != nil {
		s.userHandler.User_Route(mux, manager)
	}
	if s.productHandler != nil {
		s.productHandler.Product_Route(mux, manager)
	}

	// 4️⃣ Wrap mux with global middleware
	wrappedMux := manager.WrapMux(mux)

	// 5️⃣ Start server
	addr := ":" + strconv.Itoa(s.cnf.HttpPort)
	fmt.Printf("✅ Server running at http://%s\n", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		fmt.Println("❌ Server error:", err)
		os.Exit(1)
	}
}
