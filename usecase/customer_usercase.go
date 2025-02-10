package usecase

import (
	"f2_gc1/model"
	"f2_gc1/repository"
)

type ICustomerUseCase interface {
	Create(customer model.Customer) error
	Update(customer model.Customer) error
	Delete(id int) error
	GetAll() ([]*model.Customer, error)
	GetByID(id int) (*model.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.ICustomerRepository
}

func NewCustomerUseCase(customerRepo repository.ICustomerRepository) ICustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
	}
}

func (c *customerUseCase) Create(customer model.Customer) error {

	if customer.Name == "" && customer.Email == "" && customer.Phone == "" {
		return ErrNotFound
	}

	err := c.customerRepo.Create(&customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerUseCase) Update(customer model.Customer) error {
	if customer.ID == 0 {
		return ErrIdNotValid
	}

	err := c.customerRepo.Update(&customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerUseCase) Delete(id int) error {
	if id == 0 {
		return ErrIdNotValid
	}

	err := c.customerRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerUseCase) GetAll() ([]*model.Customer, error) {
	customers, err := c.customerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerUseCase) GetByID(id int) (*model.Customer, error) {
	customer, err := c.customerRepo.GetByID(id)
	if err != nil {
		return nil, ErrCustomerEmpty
	}
	return customer, nil
}
