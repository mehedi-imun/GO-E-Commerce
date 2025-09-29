package user

import (
	"ecommace/domain"
)

type service struct {
	UserRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		UserRepo: usrRepo,
	}
}

func (service *service) Create(user domain.User) (*domain.User, error) {
	usr, err := service.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
func (service *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := service.UserRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil

}
func (service *service) GetAll() ([]*domain.User, error) {
	usr, err := service.UserRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
