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
	GetAll(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
}
