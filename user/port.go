package user

import (
	"ecommace/domain"
	userHandler "ecommace/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
}
