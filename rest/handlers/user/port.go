package user

import "ecommace/domain"

type Service interface {
	Find(email string, pass string) (*domain.User, error)
	Create(domain.User) (*domain.User, error)
	GetAll() ([]*domain.User, error)
}
