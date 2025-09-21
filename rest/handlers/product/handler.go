package product

import (
	"ecommace/config"
)

type Handler struct {
	cnf     *config.Config
	service Service
}

func NewHandler(service Service, cnf *config.Config) *Handler {
	return &Handler{
		cnf:     cnf,
		service: service,
	}
}
