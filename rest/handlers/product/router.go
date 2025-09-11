package product

import (
	"ecommace/rest/middleware"
	"net/http"
)

func (h *Handler) Product_Route(mux *http.ServeMux, manager *middleware.Manager) {
	// router

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		))
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductByID),
		))

}
