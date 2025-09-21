package product

import (
	"net/http"

	"ecommace/config"
	"ecommace/repo"
	"ecommace/rest/middleware"
)

type Handler struct {
	mws  *middleware.Manager
	repo repo.ProductRepo
	cnf  *config.Config
}

func NewHandler(mws *middleware.Manager, repo repo.ProductRepo, cnf *config.Config) *Handler {
	return &Handler{
		mws:  mws,
		repo: repo,
		cnf:  cnf,
	}
}

// Product_Route registers all product routes
func (h *Handler) Product_Route(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /products",
		manager.With(http.HandlerFunc(h.CreateProduct)),
	)
	mux.Handle(
		"GET /products",
		manager.With(http.HandlerFunc(h.GetAllProducts)),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(http.HandlerFunc(h.GetProductByID)),
	)
}
