package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomer(conn *gorm.DB) domain.CustomerRepository {
	return &CustomerRepository{
		db: conn,
	}
}

func (cr *CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	err := cr.db.WithContext(ctx).Create(&customer).Error
	return err
}

func (cr *CustomerRepository) FindByID(ctx context.Context, id string) (customer domain.Customer, err error) {
	err = cr.db.WithContext(ctx).Take(&customer, "id = ?", id).Error

	return
}

func (cr *CustomerRepository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	err = cr.db.WithContext(ctx).Find(&customers).Error
	return
}

func (cr *CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	err := cr.db.WithContext(ctx).Save(&customer).Error
	return err
}

func (cr *CustomerRepository) Delete(ctx context.Context, id string) error {
	err := cr.db.WithContext(ctx).Delete(&domain.Customer{}, "id = ?", id).Error
	return err
}

func (cr *CustomerRepository) FindByIds(ctx context.Context, id []string) (customer []domain.Customer, err error) {
	for i := 0; i < len(id); i++ {
		var c domain.Customer
		err = cr.db.WithContext(ctx).Where("id =?", id[i]).First(&c).Error
		if err != nil {
			return nil, err
		}
		customer = append(customer, c)
	}
	return
}
