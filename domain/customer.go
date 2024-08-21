package domain

import (
	"context"
	"rest_api_sederhana/dto"
	"time"

	"gorm.io/gorm"
)

// menerapkan soft delete
type Customer struct {
	ID        string         `gorm:"primary_key;column:id"`
	Code      string         `gorm:"column:code"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	FindByID(ctx context.Context, id string) (Customer, error)
	FindAll(ctx context.Context) ([]Customer, error)
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id string) error
	FindByIds(ctx context.Context, id []string) ([]Customer, error)
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) (string, error)
	Update(ctx context.Context, req dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, id string) error
	Show(ctx context.Context, id string) (dto.CustomerData, error)
}
