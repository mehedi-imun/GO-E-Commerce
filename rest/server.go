package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"ecommace/config"

	"ecommace/rest/handlers/user"
	"ecommace/rest/middleware"
)

// Server holds dependencies
type Server struct {
	userHandler *user.Handler
	cnf         *config.Config
}

// NewServer creates a server with injected handlers and config
func NewServer(
	cnf *config.Config,

	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf: cnf,

		userHandler: userHandler,
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

	if s.userHandler != nil {
		s.userHandler.User_Route(mux, manager)
	}

	// 4️⃣ Wrap mux with global middleware
	wrappedMux := manager.WrapMux(mux)

	// 5️⃣ Start server
	addr := ":" + strconv.Itoa(s.cnf.HttpPort)
	fmt.Println("Server is running on", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		fmt.Println("Server error:", err)
	}
}
