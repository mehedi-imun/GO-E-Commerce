package cmd

import (
	"fmt"
	"os"

	"ecommace/config"
	"ecommace/infra/db"
	"ecommace/repo"
	"ecommace/rest"
	productHandler "ecommace/rest/handlers/product"
	"ecommace/user"
	userHandler "ecommace/rest/handlers/user"
	"ecommace/rest/middleware"
)

func Serve() {
	// 1️⃣ Load configuration
	cnf := config.GetConfig()

	// 2️⃣ Connect to the database
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println("❌ DB connection error:", err)
		os.Exit(1)
	}
	// Run migrations
	dbURL := db.GetMigrationDBURL(cnf.DB)
	db.RunMigrations(dbURL)

	// 3️⃣ Initialize repositories
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	// domain

	usrService:=user.NewService(userRepo)

	// 4️⃣ Initialize middleware manager
	mws := middleware.NewManager()
	mws.Use(
		middleware.Preflight,
		middleware.CORSMiddleware,
		middleware.Logger,
	)

	// 5️⃣ Initialize handlers
	userHandler := userHandler.NewHandler(cnf, usrService)
	productHandler := productHandler.NewHandler(mws, productRepo, cnf)

	// 6️⃣ Create server with config + handlers
	server := rest.NewServer(cnf, userHandler, productHandler)

	// 7️⃣ Start server
	fmt.Println("🚀 Starting server on port:", cnf.HttpPort)
	server.Start()
}
