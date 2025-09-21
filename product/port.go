package product

import (
	"ecommace/domain"
	productHandler "ecommace/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}
type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
}
