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
func (service *service) GetAll(page, limit int64) ([]*domain.Product, error) {
	p, err := service.productRepo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	return p, nil
}

func (service *service) Count() (int64, error) {
	p, err := service.productRepo.Count()
	if err != nil {
		return 0, err
	}

	return p, nil
}
