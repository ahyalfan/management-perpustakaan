package domain

import (
	"context"
	"database/sql"
	"rest_api_sederhana/dto"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `gorm:"primary_key;column:id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	Isbn        string         `gorm:"column:isbn"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeleteAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	CoverId     sql.NullString `gorm:"column:cover_id"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	Save(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
	FindByIds(ctx context.Context, id []string) ([]Book, error)
}

type BookService interface {
	Index(ctx context.Context) ([]dto.BookData, error)
	Show(ctx context.Context, id string) (dto.BookShowData, error)
	Create(ctx context.Context, req dto.CreateBookRequest) (string, error)
	Update(ctx context.Context, req dto.UpdateBookRequest) error
	Delete(ctx context.Context, id string) error
}
