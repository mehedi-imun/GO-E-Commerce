package product

import "ecommace/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	GetAll(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
}
