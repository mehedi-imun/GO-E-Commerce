package cmd

import (
	"fmt"
	"os"

	"ecommace/config"
	"ecommace/infra/db"
	"ecommace/repo"
	"ecommace/rest"
	"ecommace/rest/handlers/user"
	"ecommace/rest/middleware"
)

func Serve() {
	// 1️⃣ Load configuration
	cnf := config.GetConfig()

	// 2️⃣ Connect to the database
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("DB connection error:", err)
		os.Exit(1)
	}

	// 3️⃣ Initialize repository
	userRepo := repo.NewUserRepo(dbCon)

	// 4️⃣ Initialize middleware manager
	mws := middleware.NewManager()
	mws.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	// 5️⃣ Initialize handler
	userHandler := user.NewHandler(mws, userRepo, cnf)

	// 6️⃣ Create server with config + handler
	server := rest.NewServer(cnf, userHandler) // nil for productHandler if not needed

	// 7️⃣ Start server
	fmt.Println("Starting server on port:", cnf.HttpPort)
	server.Start()
}
