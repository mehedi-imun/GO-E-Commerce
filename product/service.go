package product

import "ecommace/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (service *service) Create(product domain.Product) (*domain.Product, error) {
	p, err := service.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	return p, nil
}
func (service *service) FindByID(id int) (*domain.Product, error) {
	p, err := service.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	return p, nil
}
func (service *service) GetAll() ([]*domain.Product, error) {
	p, err := service.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	return p, nil
}





