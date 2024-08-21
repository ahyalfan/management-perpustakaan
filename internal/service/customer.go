package service

import (
	"context"
	"errors"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(cr domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: cr,
	}
}

func (cs *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := cs.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customersData []dto.CustomerData
	for _, v := range customers {
		customersData = append(customersData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customersData, nil
}

func (cs *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) (string, error) {
	customer := domain.Customer{
		ID:   uuid.NewString(),
		Code: req.Code,
		Name: req.Name,
	}
	err := cs.customerRepository.Save(ctx, &customer)
	return customer.ID, err
}

func (cs *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	exist, err := cs.customerRepository.FindByID(ctx, req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("customer not found")
	}
	if err != nil {
		return err
	}

	exist.Code = req.Code
	exist.Name = req.Name
	return cs.customerRepository.Update(ctx, &exist)
}

func (cs *customerService) Delete(ctx context.Context, id string) error {
	_, err := cs.customerRepository.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("customer not found")
	}
	if err != nil {
		return err
	}
	print("sadsada")
	return cs.customerRepository.Delete(ctx, id)
}

func (cs *customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	customer, err := cs.customerRepository.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.CustomerData{}, errors.New("customer not found")
	}
	if err != nil {
		return dto.CustomerData{}, err
	}
	var customersData dto.CustomerData
	customersData.ID = customer.ID
	customersData.Code = customer.Code
	customersData.Name = customer.Name
	return customersData, nil
}
