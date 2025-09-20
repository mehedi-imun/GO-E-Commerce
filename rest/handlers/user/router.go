package user

import (
	"net/http"

	"ecommace/config"
	"ecommace/repo"
	"ecommace/rest/middleware"
)

type Handler struct {
	mws  *middleware.Manager
	repo repo.UserRepo
	cnf  *config.Config
}

func NewHandler(mws *middleware.Manager, repo repo.UserRepo, cnf *config.Config) *Handler {
	return &Handler{
		mws:  mws,
		repo: repo,
		cnf:  cnf,
	}
}

// User_Route registers all user routes
func (h *Handler) User_Route(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		))
	mux.Handle(
		"POST /login",
		manager.With(
			http.HandlerFunc(h.LoginUser),
		))
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(h.GetAllUsers),
		))

}
