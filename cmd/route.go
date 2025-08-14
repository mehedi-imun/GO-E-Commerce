package cmd

import (
	"ecommace/database"
	"ecommace/middleware"
	"net/http"
)

func InitRoute(mux *http.ServeMux, manager *middleware.Manager) {
	// router

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(database.ProductGet),
		))
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(database.GetProductByID),
		))

}
